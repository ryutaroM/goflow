package goflow

import "context"

type Not[T any] struct {
	predicator Predicator[T]
}

func (n Not[T]) Test(ctx context.Context, input T) (bool, error) {
	result, err := n.predicator.Test(ctx, input)
	if err != nil {
		return false, err
	}
	return !result, nil
}

func NotPredicator[T any](predicator Predicator[T]) Predicator[T] {
	return Not[T]{predicator: predicator}
}
