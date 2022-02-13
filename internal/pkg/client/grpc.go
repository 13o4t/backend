package client

import (
	"backend/internal/pkg/genprotobuf/mysvr"
	"crypto/tls"
	"crypto/x509"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewMysvrClient(addr string, useTLS bool) (client mysvr.MysvrServiceClient, close func() error, err error) {
	opts, err := grpcDialOpts(useTLS)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return mysvr.NewMysvrServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(useTLS bool) ([]grpc.DialOption, error) {
	if !useTLS {
		return []grpc.DialOption{
			//grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithInsecure(),
		}, nil
	}

	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "cannot load root CA cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    pool,
		MinVersion: tls.VersionTLS12,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}, nil
}
