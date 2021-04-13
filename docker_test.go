package supervisor

import (
	"context"
	"github.com/docker/docker/api/types"
	"testing"
	"time"
)

func TestDockerClient_ContainerList(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerConfig())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	containers, err := c.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range containers {
		t.Log(item.ID, item.Status, item.Names)
	}
}

func TestDockerClient_Apply(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerConfig())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	_, err = c.Apply(ctx, DeployConfig{
		ServicePort: 8080,
		Name:        "web",
		Image:       "juxuny/go-web",
		Tag:         "latest",
		Mounts:      nil,
		EnvData:     "",
		Envs:        nil,
	})
	if err != nil {
		t.Fatal(err)
	}
}
