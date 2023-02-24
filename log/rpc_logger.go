package log

import (
	"github.com/juxuny/env"
	log_server "github.com/juxuny/log-server"
)

var rpcLogger log_server.ClientPool

func init() {
	logServerHost := env.GetStringList("LOG_SERVER_HOST", ",")
	var err error
	if len(logServerHost) > 0 {
		rpcLogger, err = log_server.NewClientPool("", logServerHost...)
		if err != nil {
			panic(err)
		}
	}
}
