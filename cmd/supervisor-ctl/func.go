package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/juxuny/supervisor"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"strconv"
	"strings"
)

func getClient(ctx context.Context, host string) (client supervisor.SupervisorClient, err error) {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		logger.Error(err)
		return client, errors.Wrap(err, "load cert failed")
	}

	conn, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = supervisor.NewSupervisorClient(conn)
	return client, nil
}

func parseBlockSize(s string) (blockSize int64, err error) {
	s = strings.ToLower(s)
	base := int64(1)
	if strings.Contains(s, "k") {
		base *= 1 << 10
		s = strings.Replace(s, "k", "", 1)
	} else if strings.Contains(s, "m") {
		base *= 1 << 20
		s = strings.Replace(s, "m", "", 1)
	} else if strings.Contains(s, "g") {
		base *= 1 << 30
		s = strings.Replace(s, "g", "", 1)
	} else if strings.Contains(s, "t") {
		base *= 1 << 40
		s = strings.Replace(s, "t", "", 1)
	}
	blockSize, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return blockSize, err
	}
	blockSize *= base
	return
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
