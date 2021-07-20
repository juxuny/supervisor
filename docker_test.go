package supervisor

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/juxuny/supervisor/proxy"
	"strings"
	"testing"
)

func TestDockerClient_ContainerList(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerClientConfig())
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
	c, err := NewDockerClient(NewDefaultDockerClientConfig())
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
	id, err := c.Apply(ctx, DeployConfig{
		ServicePort: 8090,
		ProxyPort:   8090,
		Name:        "web",
		Image:       "juxuny/go-web",
		Tag:         "latest",
		Mounts: []*Mount{
			{HostPath: "tmp", MountPath: "/html"},
		},
		EnvData: "QUNDRVNTX0tFWT0iMTIzIDQ1NiIKU0VDUkVUPSAxMjM0NTc3OA==",
		Envs: []*KeyValue{
			{Key: "PORT", Value: "8080"},
		},
		Version: 4,
		HealthCheck: &HealthCheck{
			Type: proxy.HealthCheckType_TypeDefault,
			Path: "/",
			Port: 8090,
		},
		Restart: "always",
		Entrypoint: []string{
			"/app/go-web", "-d", "/html", "-p", "8090",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Container ID:", id)
}

func TestDockerClient_Stop(t *testing.T) {
	c, err := NewDockerClient(NewDefaultDockerClientConfig())
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

func TestParseEnv(t *testing.T) {
	s := "HUI_FU_PRI_KEY=MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMVlfYX4fTihwgAKJ2ifLN3UG9l/5cT5l9F4gP2qWJfjVSL8Pcr1vEPuF+f4S3M1gTBRSQrIkvw2QEN3+WgRfJ6DtpT7pkhGpberj37fV+r+nkVsJF/oElZbF8i26RvwM29v/lRgUPc7O/er4Py3RwB/taev+al2OXXoQ5W40+o/AgMBAAECgYA0VImIYK3hu5BQrmBwLfKZBEg1yuBA7eI/l/CqTuSZr5y8X56KFcdJQu93ga2O51pOUt5IS5Ab6M30lPO7kRc1/utjaLjB/7DarLyPpRI6Q+dJT/qepr2JKd+eGJOMDp0xX6XT278xzHhDb8DlYSUghTVwmyp31Co9uibP6iO2EQJBAPjJbt7VLjrR2obCgo2d0lPt+7k3l1pNgnF0vkSNpUs0TyA2WGBVJPSCxGDXw+2xKi3Z3M7mRirJZMzSB9MbYq0CQQDLHp9+vau8ZSvzZFWd4Kbpyf8nAW6fcpwbQG/ZCJ2o1rHKD6k6FOgPHSlkI5SGW7qGdj5y7txYZzpFInyz5MobAkANCTEACBeWCWzz5rlEhmKA91VbTShnGOye2Ukm+m0Q1brXq0FSOuPm0/tKP8QKbmARavsA9Fv03fykJtU2IJc5AkBkieDahBmYY9+QVs6GGeekeuZ/sRbHd5xLZOa336riInrYEE5sQGLo8D9HoNDofEjkO20HyLFqVJYkGEDvbkSXAkEA4tRjQvkziZf33eJAT2gICkpkgZhHXZTE8lzMoxM88knokMmj3vo1oP2jur407hTakLAHzZUgyHk3rF/DHlPgbg=="
	index := strings.Index(s, "=")
	t.Log(s[:index])
	t.Log(s[index+1:])
}
