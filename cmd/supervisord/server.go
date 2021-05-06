package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"os"
	"path"
	"time"
)

type server struct {
	supervisor.UnimplementedSupervisorServer
}

func getProxyClient(host string) (client proxy.ProxyClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = proxy.NewProxyClient(conn)
	return client, nil
}

func getSupervisorClient(host string) (client supervisor.SupervisorClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = supervisor.NewSupervisorClient(conn)
	return client, nil
}

func (t *server) ProxyStatus(ctx context.Context, req *supervisor.ProxyStatusReq) (resp *proxy.StatusResp, err error) {
	dockerClient, err := supervisor.NewDockerClient(getDockerClientConfig())
	if err != nil {
		return nil, err
	}
	proxyContainer, err := dockerClient.FindProxyContainer(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if len(proxyContainer.Ports) == 0 {
		return nil, errors.Errorf("container %s is found, but not published port", proxyContainer.ID)
	}
	p := proxyContainer.Ports[0].PublicPort
	for _, item := range proxyContainer.Ports {
		if item.PublicPort > p {
			p = item.PublicPort
		}
	}
	proxyHost := fmt.Sprintf("%s:%d", "127.0.0.1", p)
	proxyClient, err := getProxyClient(proxyHost)
	if err != nil {
		return nil, err
	}
	resp, err = proxyClient.Status(ctx, &proxy.StatusReq{})
	return resp, err
}
func (t *server) Apply(ctx context.Context, req *supervisor.ApplyReq) (resp *supervisor.ApplyResp, err error) {
	resp = &supervisor.ApplyResp{}
	dockerClient, err := supervisor.NewDockerClient(getDockerClientConfig())
	if err != nil {
		return nil, err
	}
	_, err = dockerClient.Apply(ctx, *req.Config, time.Duration(req.StopTimeout)*time.Second)
	if err != nil {
		return nil, err
	}
	if err := saveDeployConfig(*req.Config); err != nil {
		logger.Error(err)
		return nil, errors.New("save deploy config failed")
	}
	resp.Msg = "success"
	return resp, nil
}

func (t *server) Get(ctx context.Context, req *supervisor.GetReq) (resp *supervisor.GetResp, err error) {
	return
}

func (t *server) Stop(ctx context.Context, req *supervisor.StopReq) (resp *supervisor.StopResp, err error) {
	dockerClient, err := supervisor.NewDockerClient(getDockerClientConfig())
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	_, err = dockerClient.Stop(ctx, req.Name)
	return &supervisor.StopResp{}, err
}

func (t *server) Upload(ctx context.Context, req *supervisor.UploadReq) (resp *supervisor.UploadResp, err error) {
	uploadDir := path.Join(globalConfig.StoreDir, DefaultUploadDir)
	if err := touchDir(uploadDir); err != nil {
		logger.Error(err)
		return nil, err
	}
	fileName := path.Join(uploadDir, req.FileName)
	if _, err := os.Stat(fileName); os.IsExist(err) {
		if req.Force {
			err = os.Remove(fileName)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("file is exists: %s", req.FileName)
		}
	}
	mode := os.FileMode(0644)
	if req.Executable {
		mode = 0755
	}
	defer func() {
		if req.BlockNum < req.BlockNumTotal {
			return
		}
		// upload finished, compare the file hash
		var realHash string
		realHash, err = supervisor.GetFileHash(fileName, req.HashType)
		if err != nil {
			logger.Error(err)
			return
		}
		if realHash != req.FileHash {
			logger.Error("real hash: "+realHash, " need: "+req.FileHash)
			err = fmt.Errorf("incrrect file hash: %s <> %s", realHash, req.FileHash)
			return
		}
	}()
	var f *os.File
	if req.BlockNum == 1 {
		// block num = 1 need to create a new file
		logger.Info("create a new file")
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			_ = os.Remove(fileName)
		}
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, mode)
	} else {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, mode)
	}
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("write failed: %v", err)
	}
	defer f.Close()
	_, err = f.Write(req.Data)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &supervisor.UploadResp{}, nil
}
