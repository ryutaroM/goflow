package goflow

import "context"

type Flow[T any] interface {
	Process(ctx context.Context, input T) (T, error)
}

func apply[T any](ctx context.Context, it Item[T], f Flow[T]) Item[T] {
	if it.err != nil {
		return it
	}
	out, err := f.Process(ctx, it.value)
	return NewItem(ctx, out, err)
}
