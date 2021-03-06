package supervisor

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	DefaultStoreDir = "tmp"
)

type ConfigWrapper struct {
	Supervisor Config `json:"supervisor" yaml:"supervisor"`
}

type Config struct {
	ProxyImage  string `json:"proxy_image" yaml:"proxy_image"`
	ControlPort int    `json:"control_port" yaml:"control_port"`
	DockerHost  string `json:"docker_host" yaml:"docker_host"`
	StoreDir    string `json:"store_dir" yaml:"store_dir"`
	HostIp      string `json:"host_ip" yaml:"host_ip"`
	CertFile    string `json:"cert_file" yaml:"cert_file"`
	CertKeyFile string `json:"cert_key_file" yaml:"cert_key_file"`
}

func Parse(file string) (*ConfigWrapper, error) {
	config := &ConfigWrapper{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "read yaml failed")
	}
	err = yaml.Unmarshal(data, config)
	return config, err
}

func (t *Config) GetDockerClientConfig() *DockerClientConfig {
	return &DockerClientConfig{Host: t.DockerHost}
}

func Init(config Config) error {
	if config.StoreDir == "" {
		return fmt.Errorf("store dir cannot be empty")
	}
	fmt.Println("store dir:", config.StoreDir)
	if stat, err := os.Stat(config.StoreDir); os.IsNotExist(err) {
		if err := os.MkdirAll(config.StoreDir, 0776); err != nil {
			return err
		}
	} else if !stat.IsDir() {
		return fmt.Errorf("%s is not a direcory", config.StoreDir)
	}
	fmt.Println("docker host:", config.DockerHost)
	return nil
}
