package log

import (
	"github.com/juxuny/env"
	log_server "github.com/juxuny/log-server"
	"strings"
)

var rpcLogger log_server.ClientPool

func dropEmpty(list []string) []string {
	ret := make([]string, 0)
	for _, item := range list {
		if strings.TrimSpace(item) != "" {
			ret = append(ret, strings.TrimSpace(item))
		}
	}
	return ret
}

func init() {
	logServerHost := env.GetStringList("LOG_SERVER_HOST", ",")
	logServerHost = dropEmpty(logServerHost)
	var err error
	if len(logServerHost) > 0 {
		rpcLogger, err = log_server.NewClientPool("", logServerHost...)
		if err != nil {
			panic(err)
		}
	}
}
