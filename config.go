package supervisor

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigWrapper struct {
	Supervisor Config `json:"supervisor" yaml:"supervisor"`
}

type Config struct {
	ControlPort int    `json:"control_port" yaml:"control_port"`
	Docker      string `json:"docker" yaml:"docker"`
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
	return &DockerClientConfig{Host: t.Docker}
}
