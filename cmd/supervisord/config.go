package main

import "github.com/juxuny/supervisor"

const DefaultUploadDir = "uploads"

func getDockerClientConfig() supervisor.DockerClientConfig {
	config := supervisor.NewDefaultDockerClientConfig()
	if globalConfig.ProxyImage != "" {
		config.ProxyImage = globalConfig.ProxyImage
	}
	if globalConfig.DockerHost != "" {
		config.Host = globalConfig.DockerHost
	}
	return config
}
