package main

import (
	"context"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/spf13/cobra"
	"io"
	"os"
	"time"
)

var uploadFlag = struct {
	Name       string
	FilePath   string
	BlockSize  string
	Executable bool
}{}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload",
	Run: func(cmd *cobra.Command, args []string) {
		if uploadFlag.Name == "" {
			fmt.Println("name cannot be empty")
			os.Exit(-1)
		}
		if uploadFlag.FilePath == "" {
			fmt.Println("file cannot be empty")
			os.Exit(-1)
		}
		blockSize, err := parseBlockSize(uploadFlag.BlockSize)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		fileSize, err := supervisor.GetFileSize(uploadFlag.FilePath)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		blockNum := fileSize / blockSize
		if fileSize%blockSize > 0 {
			blockNum += 1
		}
		fileHash, err := supervisor.GetFileHash(uploadFlag.FilePath, supervisor.HashType_Sha256)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
		logger.Info("block num: ", blockNum, " block size:", uploadFlag.BlockSize, " file hash:", fileHash)
		for _, host := range baseFlag.Host {
			func() {

				ctx, cancel := context.WithTimeout(context.Background(), supervisor.DefaultTimeout)
				defer cancel()
				client, err := getClient(ctx, host, baseFlag.CertFile)
				if err != nil {
					logger.Error(err)
					os.Exit(-1)
				}
				f, err := os.Open(uploadFlag.FilePath)
				if err != nil {
					logger.Error(err)
					os.Exit(-1)
				}
				defer f.Close()
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
							FileName:      uploadFlag.Name,
							FileHash:      fileHash,
							HashType:      supervisor.HashType_Sha256,
							Data:          buf[:n],
							BlockNum:      uint32(index),
							BlockNumTotal: uint32(blockNum),
							FileSize:      uint64(fileSize),
							Executable:    uploadFlag.Executable,
						})
						if err != nil {
							logger.Info(fmt.Sprintf("upload(%d/%d): failed, %v", index, blockNum, err))
							os.Exit(-1)
						} else {
							logger.Info(fmt.Sprintf("upload(%d/%d): success", index, blockNum))
						}
						index += 1
					}()
				}
			}()
		}
	},
}

func init() {
	initBaseFlag(uploadCmd)
	uploadCmd.PersistentFlags().StringVar(&uploadFlag.Name, "name", "", "file name")
	uploadCmd.PersistentFlags().StringVar(&uploadFlag.FilePath, "file", "", "file to upload")
	uploadCmd.PersistentFlags().StringVar(&uploadFlag.BlockSize, "block-size", "1m", "upload block size")
	uploadCmd.PersistentFlags().BoolVar(&uploadFlag.Executable, "exec", false, "executable file")
	rootCmd.AddCommand(uploadCmd)
}
