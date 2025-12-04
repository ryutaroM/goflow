package goflow

import "context"

type Predicator[T any] interface {
	Test(ctx context.Context, input T) (bool, error)
}
