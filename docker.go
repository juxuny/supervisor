package supervisor

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"time"
)

const (
	DefaultTimeout    = time.Second * 15
	DefaultProxyImage = "juxuny/supervisor-proxy:latest"
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

func (t *DockerClient) initProxy(ctx context.Context, deployConfig DeployConfig) error {
	proxyContainerName := strings.Join([]string{containerPrefix, "proxy", deployConfig.Name}, "-")
	list, err := t.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return errors.Wrap(err, "find proxy container failed")
	}
	for _, item := range list {
		for _, n := range item.Names {
			if n == proxyContainerName {
				// proxy instance is started
				return nil
			}
		}
	}
	resp, err := t.ContainerCreate(ctx, &container.Config{
		Image: t.Config.ProxyImage,
	}, &container.HostConfig{
		AutoRemove: true,
	}, nil, nil, proxyContainerName)
	if err != nil {
		return errors.Wrap(err, "create proxy container failed")
	}
	if err := t.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return errors.Wrap(err, "start proxy container failed")
	}
	return nil
}

func (t *DockerClient) Apply(ctx context.Context, deployConfig DeployConfig) (id string, err error) {
	imageWithTag := deployConfig.Image + ":" + deployConfig.Tag
	if deployConfig.PullRetryTimes <= 0 {
		deployConfig.PullRetryTimes = 3
	}
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
	if err := t.initProxy(ctx, deployConfig); err != nil {
		return "", err
	}

	containerName := strings.Join([]string{containerPrefix, "svc", deployConfig.Name}, "-")
	resp, err := t.ContainerCreate(ctx, &container.Config{
		Image: imageWithTag,
	}, &container.HostConfig{
		AutoRemove: true,
	}, nil, nil, containerName)
	if err != nil {
		return "", errors.Wrap(err, "create container failed")
	}
	if err := t.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", errors.Wrap(err, "start container failed")
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
