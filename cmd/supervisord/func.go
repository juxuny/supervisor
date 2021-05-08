package main

import (
	"crypto/tls"
	"fmt"
	"github.com/juxuny/supervisor"
	"google.golang.org/grpc/credentials"
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
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	certFile := globalConfig.CertFile
	if certFile == "" {
		certFile = "cert/server-cert.pem"
	}
	fmt.Println("cert file:", certFile)
	certKeyFile := globalConfig.CertKeyFile
	if certKeyFile == "" {
		certKeyFile = "cert/server-key.pem"
	}
	fmt.Println("cert key file:", certKeyFile)
	serverCert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
