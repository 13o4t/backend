package app

import "backend/internal/mysvr/domain/calc"

func (a *Application) Add(numA, numB int64) int64 {
	return calc.Add(numA, numB)
}
