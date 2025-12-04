package goflow

import "context"

type Or[T any] struct {
	left  Predicator[T]
	right Predicator[T]
}

func (o Or[T]) Test(ctx context.Context, input T) (bool, error) {
	leftResult, err := o.left.Test(ctx, input)
	if err != nil {
		return false, err
	}

	if leftResult {
		return true, nil
	}

	return o.right.Test(ctx, input)
}

func OrPredicator[T any](left, right Predicator[T]) Predicator[T] {
	return Or[T]{left: left, right: right}
}