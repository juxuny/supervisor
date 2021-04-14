package supervisor

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"time"
)

const (
	DefaultTimeout    = time.Second * 30
	DefaultProxyImage = "juxuny/supervisor-proxy:latest"
	ComponentProxy    = "proxy"
	ComponentSvc      = "svc"
)

var (
	NotFound        = errors.New("not found")
	containerPrefix = "sup"
)

type DockerClientConfig struct {
	Host       string
	ProxyImage string
}

func NewDefaultDockerConfig() DockerClientConfig {
	return DockerClientConfig{
		Host:       "unix:///var/run/docker.sock",
		ProxyImage: DefaultProxyImage,
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
	return ret, NotFound
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

func (t *DockerClient) initProxy(ctx context.Context, deployConfig DeployConfig, callback func(c container.ContainerCreateCreatedBody) error) (ID string, err error) {
	proxyContainerName := t.genProxyName(deployConfig)
	// check running container
	if list, err := t.findContainer(ctx, func(container types.Container) bool {
		for _, n := range container.Names {
			if strings.Trim(n, "/") == proxyContainerName {
				return true
			}
		}
		return false
	}); err != nil {
		return "", err
	} else if len(list) > 0 {
		return list[0].ID, nil
	}
	resp, err := t.ContainerCreate(ctx, &container.Config{
		Hostname:   proxyContainerName,
		Domainname: proxyContainerName,
		Image:      t.Config.ProxyImage,
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%d", deployConfig.ServicePort)): struct{}{},
		},
	}, &container.HostConfig{
		AutoRemove: true,
		PortBindings: nat.PortMap{
			nat.Port(fmt.Sprintf("%d", deployConfig.ServicePort)): []nat.PortBinding{
				{HostPort: fmt.Sprintf("%d", deployConfig.ServicePort)},
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
	return resp.ID, nil
}

func (t *DockerClient) findContainer(ctx context.Context, filter func(container types.Container) bool) ([]types.Container, error) {
	list, err := t.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	var ret []types.Container
	for _, item := range list {
		if filter(item) {
			ret = append(ret, item)
		}
	}
	return ret, nil
}

func (t *DockerClient) genProxyName(deployConfig DeployConfig) string {
	return strings.Join([]string{containerPrefix, ComponentProxy, deployConfig.Name}, "-")
}

func (t *DockerClient) genSvcName(deployConfig DeployConfig) string {
	return strings.Join([]string{containerPrefix, ComponentSvc, deployConfig.Name, HashShort(deployConfig)}, "-")
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
	fmt.Println("pulling image:", imageWithTag)
	for i := 0; i < int(deployConfig.PullRetryTimes); i++ {
		reader, err := t.ImagePull(ctx, imageWithTag, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		_, _ = io.Copy(os.Stdout, reader)
		if _, err := t.findImage(ctx, imageWithTag); err != nil {
			if err != NotFound {
				return id, err
			}
		} else {
			break
		}
		fmt.Println("retry:", i+1)
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
	}); err != nil {
		return "", err
	} else if len(list) > 0 {
		return list[0].ID, nil
	}

	resp, err := t.ContainerCreate(ctx, &container.Config{
		Hostname:   containerName,
		Domainname: containerName,
		Image:      imageWithTag,
	}, &container.HostConfig{
		AutoRemove: true,
	}, nil, nil, containerName)
	if err != nil {
		return "", errors.Wrap(err, "create container failed")
	}
	if err := t.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", errors.Wrap(err, "start container failed")
	}
	if err := t.NetworkConnect(ctx, networkID, resp.ID, nil); err != nil {
		return resp.ID, err
	}
	//config := container.Config{
	//	Hostname:        strings.Join([]string{containerPrefix, deployConfig.Name}, "-"),
	//	Domainname:      strings.Join([]string{containerPrefix, deployConfig.Name}, "-"),
	//	User:            "root",
	//	AttachStdin:     false,
	//	AttachStdout:    false,
	//	AttachStderr:    false,
	//	ExposedPorts:    nil,
	//	Tty:             false,
	//	OpenStdin:       false,
	//	StdinOnce:       false,
	//	Env:             nil,
	//	Cmd:             nil,
	//	Healthcheck:     nil,
	//	ArgsEscaped:     false,
	//	Image:           deployConfig.Image + ":" + deployConfig.Tag,
	//	Volumes:         nil,
	//	WorkingDir:      "",
	//	Entrypoint:      nil,
	//	NetworkDisabled: false,
	//	MacAddress:      "",
	//	OnBuild:         nil,
	//	Labels:          nil,
	//	StopSignal:      "",
	//	StopTimeout:     nil,
	//	Shell:           nil,
	//}
	//hostConfig := container.HostConfig{
	//	Binds:           nil,
	//	ContainerIDFile: "",
	//	LogConfig:       container.LogConfig{},
	//	NetworkMode:     "",
	//	PortBindings:    nil,
	//	RestartPolicy:   container.RestartPolicy{},
	//	AutoRemove:      false,
	//	VolumeDriver:    "",
	//	VolumesFrom:     nil,
	//	CapAdd:          nil,
	//	CapDrop:         nil,
	//	CgroupnsMode:    "",
	//	DNS:             nil,
	//	DNSOptions:      nil,
	//	DNSSearch:       nil,
	//	ExtraHosts:      nil,
	//	GroupAdd:        nil,
	//	IpcMode:         "",
	//	Cgroup:          "",
	//	Links:           nil,
	//	OomScoreAdj:     0,
	//	PidMode:         "",
	//	Privileged:      false,
	//	PublishAllPorts: false,
	//	ReadonlyRootfs:  false,
	//	SecurityOpt:     nil,
	//	StorageOpt:      nil,
	//	Tmpfs:           nil,
	//	UTSMode:         "",
	//	UsernsMode:      "",
	//	ShmSize:         0,
	//	Sysctls:         nil,
	//	Runtime:         "",
	//	ConsoleSize:     [2]uint{},
	//	Isolation:       "",
	//	Resources:       container.Resources{},
	//	Mounts:          nil,
	//	MaskedPaths:     nil,
	//	ReadonlyPaths:   nil,
	//	Init:            nil,
	//}
	//nc := network.NetworkingConfig{}
	//t.ContainerCreate(ctx, &config, &hostConfig, &nc, nil, strings.Join([]string{containerPrefix, deployConfig.Name}, "-"))
	return
}

func (t *DockerClient) Stop(ctx context.Context, name string) (int, error) {
	proxyContainerNamePrefix := strings.Join([]string{containerPrefix, ComponentProxy, name}, "-")
	svcContainerNamePrefix := strings.Join([]string{containerPrefix, ComponentSvc, name}, "-")
	list, err := t.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return 0, err
	}
	var timeout = DefaultTimeout
	count := 0
	for _, c := range list {
		for _, n := range c.Names {
			if strings.HasPrefix(strings.Trim(n, "/"), proxyContainerNamePrefix) || strings.HasPrefix(strings.Trim(n, "/"), svcContainerNamePrefix) {
				fmt.Println("stop ", n)
				if err := t.ContainerStop(ctx, c.ID, &timeout); err != nil {
					return 0, err
				}
				count += 1
			}
		}
	}
	return count, nil
}
