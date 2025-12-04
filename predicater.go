package goflow

import "context"

type Predicater[T any] interface {
	Test(ctx context.Context, input T) (bool, error)
}
