package supervisor

import (
	"crypto/sha256"
	"fmt"
	pb "github.com/juxuny/supervisor/proxy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"math/rand"
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
		h.Write([]byte(fmt.Sprintf("%s:%v\n", ft.Name, fv.Interface())))
	}
	//fmt.Println(tt.Kind(), vv.NumField())
	return fmt.Sprintf("%X", h.Sum(nil))
}

func HashShort(v interface{}) string {
	h := Hash(v)
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

func init() {
	rand.Seed(time.Now().UnixNano())
}
