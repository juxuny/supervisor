package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var uploadMultipleFlag = struct {
	FileList   []string
	BlockSize  string
	Executable bool
}{}

func uploadFile(local, remoteName string, blockSize int64) (err error) {
	fileSize, err := supervisor.GetFileSize(local)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	blockNum := fileSize / blockSize
	if fileSize%blockSize > 0 {
		blockNum += 1
	}
	fileHash, err := supervisor.GetFileHash(local, supervisor.HashType_Sha256)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	logger.Info("block num: ", blockNum, " block size:", uploadMultipleFlag.BlockSize, " file hash:", fileHash)
	ctx, cancel := context.WithTimeout(context.Background(), supervisor.DefaultTimeout)
	defer cancel()
	client, err := getClient(ctx, baseFlag.Host, baseFlag.CertFile)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	f, err := os.Open(local)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	buf := make([]byte, blockSize)
	uploading := true
	index := 1
	for uploading {
		func() {
			n, err := f.Read(buf)
			if err == io.EOF {
				uploading = false
				return
			}
			uploadCtx, uploadCancel := context.WithTimeout(context.Background(), time.Duration(baseFlag.Timeout)*time.Second)
			defer uploadCancel()
			_, err = client.Upload(uploadCtx, &supervisor.UploadReq{
				FileName:      remoteName,
				FileHash:      fileHash,
				HashType:      supervisor.HashType_Sha256,
				Data:          buf[:n],
				BlockNum:      uint32(index),
				BlockNumTotal: uint32(blockNum),
				FileSize:      uint64(fileSize),
				Executable:    uploadMultipleFlag.Executable,
			})
			if err != nil {
				logger.Info(fmt.Sprintf("%s upload(%d/%d): failed, %v", local, index, blockNum, err))
				os.Exit(-1)
			} else {
				logger.Info(fmt.Sprintf("%s upload(%d/%d): success", local, index, blockNum))
			}
			index += 1
		}()
	}
	return nil
}

var uploadMultipleCmd = &cobra.Command{
	Use:   "upload-multiple",
	Short: "upm",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("init upload")
		m := make(map[string]string)
		if len(uploadMultipleFlag.FileList) == 0 {
			Fatal("file cannot be empty")
		}
		for _, item := range uploadMultipleFlag.FileList {
			l := strings.Split(item, ":")
			if len(l) == 2 {
				m[l[0]] = l[1]
				logger.Info("upload file " + l[0] + " to " + l[1])
			} else {
				Fatal("invalid upload file: ", item)
			}
		}
		blockSize, err := parseBlockSize(uploadMultipleFlag.BlockSize)
		if err != nil {
			Fatal(err)
		}
		wg := sync.WaitGroup{}
		for local, remote := range m {
			wg.Add(1)
			go func(local, remoteName string, blockSize int64) {
				defer func() {
					wg.Done()
				}()
				if err := uploadFile(local, remoteName, blockSize); err != nil {
					logger.Error(err)
					return
				}
			}(local, remote, blockSize)
		}
		wg.Wait()
		logger.Info("upload finished")
	},
}

func init() {
	initBaseFlag(uploadMultipleCmd)
	uploadMultipleCmd.PersistentFlags().StringSliceVar(&uploadMultipleFlag.FileList, "file", nil, "file name, e.g /local/file/path:{REMOTE_FILE_NAME}")
	uploadMultipleCmd.PersistentFlags().StringVar(&uploadMultipleFlag.BlockSize, "block-size", "1m", "uploadMultiple block size")
	uploadMultipleCmd.PersistentFlags().BoolVar(&uploadMultipleFlag.Executable, "exec", false, "executable file")
	rootCmd.AddCommand(uploadMultipleCmd)
}
