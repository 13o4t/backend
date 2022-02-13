package ports

import (
	"backend/internal/mysvr/app"
	"backend/internal/pkg/genprotobuf/mysvr"
	"context"
)

type GRPCServer struct {
	app app.Application
	mysvr.UnimplementedMysvrServiceServer
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (s *GRPCServer) Add(ctx context.Context, request *mysvr.AddRequest) (*mysvr.AddResponse, error) {
	a := request.A
	b := request.B

	result := s.app.Add(a, b)

	response := &mysvr.AddResponse{
		Result: result,
	}
	return response, nil
}
