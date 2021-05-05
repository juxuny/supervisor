package supervisor

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/juxuny/env"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"hash"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

func Hash(v interface{}) string {
	tt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	if tt.Kind() == reflect.Ptr {
		tt = reflect.TypeOf(vv.Elem().Interface())
		vv = reflect.ValueOf(vv.Elem().Interface())
	}
	h := sha256.New()
	for i := 0; i < vv.NumField(); i++ {
		if strings.HasPrefix(tt.Field(i).Name, "XXX_") {
			continue
		}
		ft := tt.Field(i)
		fv := vv.Field(i)
		//fmt.Println("hash: ", fv.Interface())
		h.Write([]byte(fmt.Sprintf("%s:%v\n", ft.Name, fv.Interface())))
	}
	//fmt.Println(tt.Kind(), vv.NumField())
	return fmt.Sprintf("%X", h.Sum(nil))
}

func HashShort(v interface{}) string {
	h := Hash(v)
	fmt.Println("hash:", h)
	if len(h) > 10 {
		return h[:10]
	}
	return h
}

func createProxyControlClient(host string) (client pb.ProxyClient, err error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = pb.NewProxyClient(conn)
	return client, nil
}

func randNum(min, max int) int {
	n := rand.Intn(max - min)
	return n + min
}

func GetFileSize(fileName string) (size int64, err error) {
	stat, err := os.Stat(fileName)
	if err != nil {
		return size, err
	}
	size = stat.Size()
	return
}

func getWd() (string, error) {
	pwd := env.GetString("HOST_PWD", "")
	if pwd == "" {
		return os.Getwd()
	}
	return pwd, nil
}

func GetFileHash(fileName string, hashType HashType) (ret string, err error) {
	var h hash.Hash
	switch hashType {
	case HashType_MD5:
		h = md5.New()
	case HashType_Sha1:
		h = sha1.New()
	case HashType_Sha256:
		h = sha256.New()
	default:
		return "", errors.New("unknown hash type: " + fmt.Sprintf("%v", hashType))
	}
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf := make([]byte, DefaultBlockSize)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		h.Write(buf[:n])
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
