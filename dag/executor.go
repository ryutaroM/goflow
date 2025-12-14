package dag

import (
	"context"
)

type Result[T any] struct {
	values map[string]T
}

type Executor[T any] interface {
	Execute(ctx context.Context, in Result[T]) (Result[T], error)
}
