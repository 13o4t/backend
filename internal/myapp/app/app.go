package app

import (
	"backend/internal/mysvr/adaptors"
	"backend/internal/pkg/client"
	"context"
)

type Application struct {
	mysvrService MysvrService
}

type Options struct {
	GRPCUseTLS bool

	MysvrGRPCAddr string
}

func NewApplication(ctx context.Context, opt Options) (app Application, cleanup func(), err error) {
	mysvrClient, closeMysvrClient, err := client.NewMysvrClient(opt.MysvrGRPCAddr, opt.GRPCUseTLS)
	if err != nil {
		return Application{}, func() {}, err
	}

	mysvrGRPC := adaptors.NewMysvrGRPC(mysvrClient)

	app = Application{
		mysvrService: mysvrGRPC,
	}

	cleanup = func() {
		_ = closeMysvrClient()
	}

	return app, cleanup, nil
}
