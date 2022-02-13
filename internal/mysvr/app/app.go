package app

import "context"

type Application struct {
}

func NewApplication(ctx context.Context) Application {
	return Application{}
}
