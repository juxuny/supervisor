package supervisor

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

const (
	DefaultTimeout    = time.Second * 60
	DefaultBlockSize  = 1 << 20 // 1M
	DefaultProxyImage = "juxuny/supervisor-proxy:latest"
	ComponentProxy    = "proxy"
	ComponentSvc      = "svc"
	ControlPortOffset = 100
)

const (
	ContainerStatusRunning = "running"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrHealthCheckTimeout = errors.New("health check timeout")
	containerPrefix       = "sup"
	DefaultTimeoutPointer *time.Duration
)

func init() {
	tmp := DefaultTimeout
	DefaultTimeoutPointer = &tmp
}

type DockerClientConfig struct {
	Host       string
	ProxyImage string
	HostIp     string // supervisord bind ip address
}

func NewDefaultDockerClientConfig() DockerClientConfig {
	return DockerClientConfig{
		Host:       "unix:///var/run/docker.sock",
		ProxyImage: DefaultProxyImage,
		HostIp:     "",
	}
}

type DockerClient struct {
	Config DockerClientConfig
	*client.Client
}

func NewDockerClient(config DockerClientConfig) (*DockerClient, error) {
	sdk := &DockerClient{
		Config: config,
	}
	var err error
	sdk.Client, err = client.NewClientWithOpts(func(c *client.Client) error {
		if config.Host != "" {
			if err := client.WithHost(config.Host)(c); err != nil {
				return err
			}
		} else if host := os.Getenv("DOCKER_HOST"); host != "" {
			if err := client.WithHost(host)(c); err != nil {
				return err
			}
		}

		if version := os.Getenv("DOCKER_API_VERSION"); version != "" {
			if err := client.WithVersion(version)(c); err != nil {
				return err
			}
		}
		return nil
	})
	return sdk, err
}

func (t *DockerClient) findImage(ctx context.Context, imageWithTag string) (ret types.ImageSummary, err error) {
	list, err := t.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return ret, errors.Wrap(err, "get image list failed")
	}
	for _, item := range list {
		for _, repoTag := range item.RepoTags {
			if repoTag == imageWithTag {
				return item, nil
			}
		}
	}
	return ret, ErrNotFound
}

func (t *DockerClient) findNetwork(ctx context.Context, filter func(item types.NetworkResource) bool) (ret types.NetworkResource, found bool, err error) {
	list, err := t.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		return ret, false, err
	}
	for _, item := range list {
		if filter(item) {
			return item, true, nil
		}
	}
	return ret, false, nil
}

func (t *DockerClient) initNetwork(ctx context.Context, deployConfig DeployConfig) (id string, err error) {
	networkName := strings.Join([]string{containerPrefix, "network", deployConfig.Name}, "-")
	fmt.Println("init network:", networkName)
	res, found, err := t.findNetwork(ctx, func(item types.NetworkResource) bool {
		return item.Name == networkName
	})
	if err != nil {
		return "", err
	}
	if found {
		return res.ID, nil
	}
	fmt.Println(res.ID)
	resp, err := t.NetworkCreate(ctx, networkName, types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         "",
		Scope:          "",
		IPAM:           nil,
		Internal:       false,
		Attachable:     false,
		Labels:         nil,
	})
	if err != nil {
		return id, err
	}
	return resp.ID, nil
}

func (t *DockerClient) createProxyEnv(deployConfig DeployConfig) []string {
	svcContainerName := t.genSvcName(deployConfig)
	return []string{
		"REMOTE=" + svcContainerName + fmt.Sprintf(":%d", deployConfig.ServicePort),
		"CONTROL_PORT=" + fmt.Sprintf("%d", deployConfig.ProxyPort+ControlPortOffset),
		"LISTEN_PORT=" + fmt.Sprintf("%d", deployConfig.ProxyPort),
	}
}

func (t *DockerClient) initProxy(ctx context.Context, deployConfig DeployConfig, callback func(c container.ContainerCreateCreatedBody) error) (ID string, err error) {
	fmt.Println("init proxy image: ", t.Config.ProxyImage)
	err = t.initImage(ctx, deployConfig, t.Config.ProxyImage)
	if err != nil {
		return "", err
	}
	proxyContainerName := t.genProxyName(deployConfig)
	// check running container
	if list, err := t.findContainer(ctx, func(container types.Container) bool {
		for _, n := range container.Names {
			if strings.Trim(n, "/") == proxyContainerName {
				return true
			}
		}
		return false
	}, true); err != nil {
		return "", err
	} else if len(list) > 0 {
		if list[0].State == ContainerStatusRunning {
			return list[0].ID, nil
		}
		if err := t.ContainerRemove(ctx, list[0].ID, types.ContainerRemoveOptions{}); err != nil {
			fmt.Printf("auto clean up proxy container '%s' failed\n", list[0].ID)
			return "", err
		}
	}
	resp, err := t.ContainerCreate(ctx, &container.Config{
		Hostname:   proxyContainerName,
		Domainname: proxyContainerName,
		Image:      t.Config.ProxyImage,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%d", deployConfig.ProxyPort)):                   struct{}{},
			nat.Port(fmt.Sprintf("%d", deployConfig.ProxyPort+ControlPortOffset)): struct{}{},
		},
		Env: t.createProxyEnv(deployConfig),
	}, &container.HostConfig{
		AutoRemove: true,
		PortBindings: nat.PortMap{
			nat.Port(fmt.Sprintf("%d", deployConfig.ProxyPort)): []nat.PortBinding{
				{HostPort: fmt.Sprintf("%d", deployConfig.ProxyPort)},
			},
			nat.Port(fmt.Sprintf("%d", deployConfig.ProxyPort+ControlPortOffset)): []nat.PortBinding{
				{HostPort: fmt.Sprintf("%d", deployConfig.ProxyPort+ControlPortOffset)},
			},
		},
	}, nil, nil, proxyContainerName)
	if err != nil {
		return "", errors.Wrap(err, "create proxy container failed")
	}
	if err := t.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", errors.Wrap(err, "start proxy container failed")
	}
	if err := callback(resp); err != nil {
		return resp.ID, err
	}
	fmt.Println("started:", resp.ID)
	return resp.ID, nil
}

func (t *DockerClient) findContainer(ctx context.Context, filter func(container types.Container) bool, all bool) ([]types.Container, error) {
	defaultFilter := func(c types.Container) bool {
		for _, n := range c.Names {
			if strings.HasPrefix(strings.Trim(n, "/"), containerPrefix) {
				return true
			}
		}
		return false
	}
	list, err := t.ContainerList(ctx, types.ContainerListOptions{All: all})
	if err != nil {
		return nil, err
	}
	var ret []types.Container
	for _, item := range list {
		if defaultFilter(item) && filter(item) {
			ret = append(ret, item)
		}
	}
	return ret, nil
}

func (t *DockerClient) genProxyName(deployConfig DeployConfig) string {
	return t.genProxyNameByServiceName(deployConfig.Name)
}

func (t *DockerClient) genProxyNameByServiceName(name string) string {
	return strings.Join([]string{containerPrefix, ComponentProxy, name}, "-")
}

func (t *DockerClient) genSvcName(deployConfig DeployConfig) string {
	return strings.Join([]string{containerPrefix, ComponentSvc, deployConfig.Name, HashShort(deployConfig)}, "-")
}

func (t *DockerClient) genSvcNameWithoutHash(deployConfig DeployConfig) string {
	return strings.Join([]string{containerPrefix, ComponentSvc, deployConfig.Name}, "-")
}

func (t *DockerClient) initImage(ctx context.Context, deployConfig DeployConfig, imageWithTag string) error {
	fmt.Println("check image:", imageWithTag)
	_, err := t.findImage(ctx, imageWithTag)
	if err != nil {
		if err != ErrNotFound {
			return errors.Wrap(err, "find image failed,"+imageWithTag)
		}
	} else {
		return nil // image is exists
	}
	for i := 0; i < int(deployConfig.PullRetryTimes); i++ {
		reader, err := t.ImagePull(ctx, imageWithTag, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		_, _ = io.Copy(os.Stdout, reader)
		if _, err := t.findImage(ctx, imageWithTag); err != nil {
			if err != ErrNotFound {
				return err
			}
		} else {
			break
		}
		fmt.Println("retry:", i+1)
	}
	return nil
}

func (t *DockerClient) Apply(ctx context.Context, deployConfig DeployConfig) (id string, err error) {
	imageWithTag := deployConfig.Image + ":" + deployConfig.Tag
	if deployConfig.PullRetryTimes <= 0 {
		deployConfig.PullRetryTimes = 3
	}
	networkID, err := t.initNetwork(ctx, deployConfig)
	if err != nil {
		return id, err
	}
	err = t.initImage(ctx, deployConfig, imageWithTag)
	if err != nil {
		return id, err
	}
	_, err = t.initProxy(ctx, deployConfig, func(c container.ContainerCreateCreatedBody) error {
		return t.NetworkConnect(ctx, networkID, c.ID, nil)
	})
	if err != nil {
		return "", err
	}

	containerName := t.genSvcName(deployConfig)
	// check running container
	if list, err := t.findContainer(ctx, func(container types.Container) bool {
		for _, n := range container.Names {
			if strings.Trim(n, "/") == containerName {
				return true
			}
		}
		return false
	}, true); err != nil {
		return "", err
	} else if len(list) > 0 {
		fmt.Println("container ", containerName, " is ", list[0].State)
		if list[0].State == ContainerStatusRunning {
			return list[0].ID, nil
		}
		if err := t.ContainerRemove(ctx, list[0].ID, types.ContainerRemoveOptions{}); err != nil {
			fmt.Printf("[check running container]auto remove container '%s' failed\n", list[0].ID)
			return "", err
		}
	}
	envs, err := t.parseEnv(deployConfig)
	if err != nil {
		return "", err
	}
	fmt.Println("creating svc")
	resp, err := t.ContainerCreate(ctx, &container.Config{
		Hostname:   containerName,
		Domainname: containerName,
		Image:      imageWithTag,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%d", deployConfig.ServicePort)): struct{}{},
		},
		Env:        envs,
		Entrypoint: deployConfig.Entrypoint,
	}, &container.HostConfig{
		AutoRemove: deployConfig.Restart == "" || deployConfig.Restart == "no",
		//PortBindings: nat.PortMap{
		//	nat.Port(fmt.Sprintf("%d", deployConfig.ServicePort)): []nat.PortBinding{
		//		{HostPort: fmt.Sprintf("%d", deployConfig.ServicePort + uint32(randNum(1000, 2000)))},
		//	},
		//},
		RestartPolicy: container.RestartPolicy{Name: deployConfig.Restart},
		Mounts:        t.parseMounts(deployConfig),
	}, nil, nil, containerName)
	if err != nil {
		return "", errors.Wrap(err, "create container failed")
	}
	fmt.Println("created svc: ", resp.ID)
	fmt.Println("starting svc")
	if err := t.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", errors.Wrap(err, "start container failed")
	}
	fmt.Println("started svc:", resp.ID)
	fmt.Println("bind network")
	if err := t.NetworkConnect(ctx, networkID, resp.ID, nil); err != nil {
		return resp.ID, err
	}

	fmt.Println("creating proxy control client: ", t.Config.HostIp)
	proxyClient, err := createProxyControlClient(fmt.Sprintf("%s:%d", t.Config.HostIp, deployConfig.ProxyPort+ControlPortOffset))
	if err != nil {
		return "", err
	}
	if err := t.waitingHealthCheck(ctx, proxyClient, deployConfig); err != nil {
		return "", err
	}
	fmt.Println("update proxy svc")
	_, err = proxyClient.Update(ctx, &proxy.UpdateReq{
		Status: &proxy.Status{
			Remote: fmt.Sprintf("%s:%d", containerName, deployConfig.ServicePort),
		},
	})
	if err != nil {
		return "", err
	}

	fmt.Println("clean up old instances")
	runningContainerList, err := t.findContainer(ctx, func(container types.Container) bool {
		for _, n := range container.Names {
			if strings.Trim(n, "/") == containerName || !strings.HasPrefix(n, "/"+t.genSvcNameWithoutHash(deployConfig)) {
				return false
			}
		}
		return true
	}, false)
	if len(runningContainerList) == 0 {
		fmt.Println("running container count: ", len(runningContainerList))
		return resp.ID, nil
	}
	if err := t.stopRunningContainer(ctx, runningContainerList...); err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp.ID, nil
}

func (t *DockerClient) stopRunningContainer(ctx context.Context, c ...types.Container) error {
	for _, item := range c {
		fmt.Println("stopping container: ", item.ID, item.Names)
		if err := t.ContainerStop(ctx, item.ID, DefaultTimeoutPointer); err != nil {
			return err
		}
	}
	return nil
}

func (t *DockerClient) Stop(ctx context.Context, name string) (int, error) {
	proxyContainerNamePrefix := strings.Join([]string{containerPrefix, ComponentProxy, name}, "-")
	svcContainerNamePrefix := strings.Join([]string{containerPrefix, ComponentSvc, name}, "-")
	list, err := t.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return 0, err
	}
	count := 0
	for _, c := range list {
		for _, n := range c.Names {
			if strings.HasPrefix(strings.Trim(n, "/"), proxyContainerNamePrefix) || strings.HasPrefix(strings.Trim(n, "/"), svcContainerNamePrefix) {
				fmt.Println("stop ", n)
				if err := t.ContainerStop(ctx, c.ID, DefaultTimeoutPointer); err != nil {
					return 0, err
				}
				count += 1
			}
		}
	}
	return count, nil
}

func (t *DockerClient) parseEnv(deployConfig DeployConfig) ([]string, error) {
	var ret []string
	m := make(map[string]string)
	if deployConfig.EnvData != "" {
		data, err := base64.StdEncoding.DecodeString(deployConfig.EnvData)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			l := strings.Trim(line, " ")
			if strings.HasPrefix(l, "#") {
				continue
			}
			index := strings.Index(l, "=")
			if index < 0 {
				continue
			}
			k := l[:index]
			v := l[index:]
			k = strings.Trim(k, "\" =")
			v = strings.Trim(v, "=\n\" '")
			m[k] = v
		}
	}
	for _, kv := range deployConfig.Envs {
		m[kv.Key] = kv.Value
	}
	for k, v := range m {
		ret = append(ret, k+"="+v)
	}
	return ret, nil
}

func (t *DockerClient) parseMounts(deployConfig DeployConfig) []mount.Mount {
	ret := make([]mount.Mount, 0)
	if deployConfig.Mounts == nil {
		return ret
	}
	for _, m := range deployConfig.Mounts {
		hostPath := m.HostPath
		if strings.HasPrefix(hostPath, "./") || !strings.HasPrefix(hostPath, "/") {
			wd, err := getWd()
			if err != nil {
				fmt.Println(err)
				continue
			}
			hostPath = path.Join(wd, strings.Replace(hostPath, "./", "", 1))
		}
		ret = append(ret, mount.Mount{
			Type:   mount.TypeBind,
			Source: hostPath,
			Target: m.MountPath,
		})
	}
	return ret
}

func (t *DockerClient) HealthCheck(ctx context.Context, client proxy.ProxyClient, deployConfig DeployConfig) error {
	if deployConfig.HealthCheck == nil {
		return fmt.Errorf("health check config is empty(nil)")
	}
	svcContainerName := t.genSvcName(deployConfig)
	_, err := client.Check(ctx, &proxy.CheckReq{
		Type: deployConfig.HealthCheck.Type,
		Host: svcContainerName,
		Path: deployConfig.HealthCheck.Path,
		Port: deployConfig.HealthCheck.Port,
	})
	return err
}

func (t *DockerClient) waitingHealthCheck(ctx context.Context, client proxy.ProxyClient, config DeployConfig) error {
	fmt.Println("waiting health check")
	count := 0
	for {
		select {
		case <-ctx.Done():
			return ErrHealthCheckTimeout
		default:
		}
		if err := t.HealthCheck(ctx, client, config); err != nil {
			count += 1
			fmt.Println(err)
			fmt.Println("health check retry in 3 seconds, times:", count)
			time.Sleep(time.Second * 3)
			continue
		} else {
			break
		}
	}
	return nil
}

func (t *DockerClient) FindProxyContainer(ctx context.Context, name string) (ret *types.Container, err error) {
	proxyContainerName := t.genProxyNameByServiceName(name)
	list, err := t.findContainer(ctx, func(container types.Container) bool {
		for _, n := range container.Names {
			s := strings.Trim(n, "/")
			if s == proxyContainerName {
				return true
			}
		}
		return false
	}, true)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, ErrNotFound
	}
	if list[0].State == ContainerStatusRunning {
		return &list[0], nil
	}
	if err := t.ContainerRemove(ctx, list[0].ID, types.ContainerRemoveOptions{}); err != nil {
		fmt.Printf("auto clean up exited container '%s' failed\n", list[0].ID)
	}
	return nil, ErrNotFound
}
