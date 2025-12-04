package goflow

import "context"

type Transformer[In, Out any] interface {
	Transform(ctx context.Context, input In) (Out, error)
}
