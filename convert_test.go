package supervisor

import (
	"encoding/json"
	"testing"
)

func TestDecodeYAML(t *testing.T) {
	dc, err := LoadDeployConfigFromYaml("tmp/deploy-web.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dc)
	jsonData, _ := json.Marshal(dc)
	t.Log(string(jsonData))
}
