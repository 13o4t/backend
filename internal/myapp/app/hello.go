package app

import (
	"backend/internal/myapp/domain/hello"
	"context"
)

func (a *Application) Hello(ctx context.Context, name string) string {
	return hello.Hello(name)
}
