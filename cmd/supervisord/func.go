package main

import (
	"fmt"
	"github.com/juxuny/supervisor"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func saveDeployConfig(deployConfig supervisor.DeployConfig) error {
	h := supervisor.HashShort(deployConfig)
	logger.Info("save deploy config: ", h)
	storeDir := globalConfig.StoreDir
	if globalConfig.StoreDir == "" {
		logger.Info("used default store dir:", supervisor.DefaultStoreDir)
		storeDir = supervisor.DefaultStoreDir
	}
	fileName := path.Join(storeDir, strings.Join([]string{deployConfig.Name, h}, "-")+".yaml")
	data, err := yaml.Marshal(deployConfig)
	if err != nil {
		logger.Error(err)
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func touchDir(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0776)
	}
	if stat.IsDir() {
		return nil
	} else {
		return fmt.Errorf("path %s is not a director", dir)
	}
	return nil
}
