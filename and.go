package goflow

import "context"

type And[T any] struct {
	left  Predicator[T]
	right Predicator[T]
}

func (a And[T]) Test(ctx context.Context, input T) (bool, error) {
	leftResult, err := a.left.Test(ctx, input)
	if err != nil {
		return false, err
	}

	if !leftResult {
		return false, nil
	}

	return a.right.Test(ctx, input)
}

func AndPredicator[T any](left, right Predicator[T]) Predicator[T] {
	return And[T]{left: left, right: right}
}
