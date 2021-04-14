package supervisor

import "testing"

func TestHash(t *testing.T) {
	dc := DeployConfig{
		ServicePort: 8080,
		Name:        "web",
		Image:       "juxuny/go-web",
		Tag:         "latest",
		Mounts: []*Mount{
			{HostPath: "/var/www", MountPath: "/var/www"},
			{HostPath: "/mnt", MountPath: "/mnt"},
		},
		EnvData: "",
		Envs: []*KeyValue{
			{Key: "CONTROL_PORT", Value: "8080"},
			{Key: "LISTEN_PORT", Value: "8888"},
		},
	}
	s := Hash(&dc)
	t.Log(s)
	for i := 0; i < 100; i++ {
		if Hash(dc) != s {
			t.Fatal("wrong hash")
		}
	}
	t.Log(HashShort(dc))
}
