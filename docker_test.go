package supervisor

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/juxuny/supervisor/proxy"
	"testing"
)

func TestDockerClient_ContainerList(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerConfig())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
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
	//if num, err := c.Stop(ctx, "web"); err != nil {
	//	t.Fatal(err)
	//} else {
	//	t.Log("stop containers:", num)
	//}
	id, err := c.Apply(ctx, DeployConfig{
		ServicePort: 8080,
		Name:        "web",
		Image:       "juxuny/go-web",
		Tag:         "latest",
		Mounts:      nil,
		EnvData:     "",
		Envs: []*KeyValue{
			{Key: "PORT", Value: "8080"},
		},
		Version: 12,
		HealthCheck: &HealthCheck{
			Type: proxy.HealthCheckType_TypeDefault,
			Path: "/",
			Port: 8080,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Container ID:", id)
}

func TestDockerClient_Stop(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerConfig())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	if num, err := c.Stop(ctx, "web"); err != nil {
		t.Fatal(err)
	} else {
		t.Log("stop containers:", num)
	}
}
