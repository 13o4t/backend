package app

import "context"

func (a *Application) Add(ctx context.Context, numA, numB int64) (int64, error) {
	return a.mysvrService.Add(ctx, numA, numB)
}
