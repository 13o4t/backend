package app

import "context"

type MysvrService interface {
	Add(ctx context.Context, a, b int64) (int64, error)
}
