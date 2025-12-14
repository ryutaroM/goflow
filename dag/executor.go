package dag

import "context"

type Executor[T any] interface {
	Execute(ctx context.Context, in Data[T]) (Data[T], error)
}
