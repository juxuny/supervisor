package proxy

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Proxy struct {
	ControlPort uint32 `json:"control_port" yaml:"control_port"`
	ListenPort  uint32 `json:"listen_port" yaml:"listen_port"`
	Remote      string `json:"remote" yaml:"remote"`
	ReadTimeout int    `json:"read_timeout" yaml:"read_timeout"`
}

type Config struct {
	Proxy Proxy `json:"proxy" yaml:"proxy"`
}

func Parse(file string) (config *Config, err error) {
	config = &Config{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "read yaml failed")
	}
	err = yaml.Unmarshal(data, config)
	return
}
