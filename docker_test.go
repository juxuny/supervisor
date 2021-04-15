package supervisor

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/juxuny/supervisor/proxy"
	"os"
	"path"
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
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if num, err := c.Stop(ctx, "web"); err != nil {
		t.Fatal(err)
	} else {
		t.Log("stop containers:", num)
	}
	id, err := c.Apply(ctx, DeployConfig{
		ServicePort: 8080,
		ProxyPort:   8090,
		Name:        "web",
		Image:       "juxuny/go-web",
		Tag:         "latest",
		Mounts: []*Mount{
			{HostPath: path.Join(wd, "tmp"), MountPath: "/html"},
		},
		EnvData: "QUNDRVNTX0tFWT0iMTIzIDQ1NiIKU0VDUkVUPSAxMjM0NTc3OA==",
		Envs: []*KeyValue{
			{Key: "PORT", Value: "8080"},
		},
		Version: 3,
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
