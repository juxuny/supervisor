package supervisor

import (
	"encoding/json"
	//"gopkg.in/yaml.v2"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func LoadDeployConfigFromYaml(fileName string) (deployConfig DeployConfig, err error) {
	data, err := ioutil.ReadFile(fileName)
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return deployConfig, err
	}
	var wrapper struct {
		Deploy DeployConfig `json:"deploy" yaml:"deploy"`
	}
	err = json.Unmarshal(jsonData, &wrapper)
	return wrapper.Deploy, err
}
