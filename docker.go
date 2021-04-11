package supervisor

type DockerClientConfig struct {
	Host string
}

type DockerClient struct {
	Config DockerClientConfig
}

func NewDockerClient(config DockerClientConfig) (*DockerClient, error) {
	client := &DockerClient{
		Config: config,
	}
	return client, nil
}
