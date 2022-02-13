package adaptors

import (
	"backend/internal/pkg/genprotobuf/mysvr"
	"context"
)

type MysvrGRPC struct {
	client mysvr.MysvrServiceClient
}

func NewMysvrGRPC(c mysvr.MysvrServiceClient) MysvrGRPC {
	return MysvrGRPC{client: c}
}

func (s MysvrGRPC) Add(ctx context.Context, a, b int64) (int64, error) {
	resp, err := s.client.Add(ctx, &mysvr.AddRequest{
		A: a,
		B: b,
	})

	if err != nil {
		return -1, err
	}

	return resp.Result, nil
}
