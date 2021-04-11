package supervisor

import (
	"github.com/docker/docker/client"
	"os"
)

type DockerClientConfig struct {
	Host string
}

func NewDefaultDockerConfig() DockerClientConfig {
	return DockerClientConfig{Host: "unix:///var/run/docker.sock"}
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
