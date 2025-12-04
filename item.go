package goflow

import "context"

type Item[T any] struct {
	ctx   context.Context
	value T
	err   error
}

func (it Item[T]) Value() T {
	return it.value
}

func (it Item[T]) Error() error {
	return it.err
}

func (it Item[T]) Result() (T, error) {
	return it.value, it.err
}

func NewItem[T any](ctx context.Context, v T, err error) Item[T] {
	return Item[T]{ctx: ctx, value: v, err: err}
}

func (it Item[T]) Pipe(f Flow[T]) Item[T] {
	return apply(it.ctx, it, f)
}

func (it Item[T]) Branch(pd Predicator[T], left, right Flow[T]) Item[T] {
	if it.err != nil {
		return it
	}

	result, err := pd.Test(it.ctx, it.value)
	if err != nil {
		var zero T
		return NewItem(it.ctx, zero, err)
	}

	if result {
		return it.Pipe(left)
	}
	return it.Pipe(right)
}

func Transform[IN, OUT any](it Item[IN], t Transformer[IN, OUT]) Item[OUT] {
	if it.err != nil {
		var zero OUT
		return NewItem(it.ctx, zero, it.err)
	}
	output, err := t.Transform(it.ctx, it.value)
	return NewItem(it.ctx, output, err)
}
