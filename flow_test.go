package goflow

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type DoubleFlow struct{}

func (f DoubleFlow) Process(ctx context.Context, input int) (int, error) {
	return input * 2, nil
}

type ErrorFlow struct{}

func (f ErrorFlow) Process(ctx context.Context, input int) (int, error) {
	return 0, errors.New("ErrFlowProcessing")
}

func TestApply(t *testing.T) {
	type test struct {
		name     string
		item     Item[int]
		flow     Flow[int]
		expected Item[int]
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			result := apply(tt.item.ctx, tt.item, tt.flow)
			if !reflect.DeepEqual(result.value, tt.expected.value) {
				t.Fatalf("got %v, want %v", result.value, tt.expected.value)
			}

			if result.err != nil || tt.expected.err != nil {
				if result.err == nil || tt.expected.err == nil || result.err.Error() != tt.expected.err.Error() {
					t.Fatalf("error mismatch: got %v, want %v", result.err, tt.expected.err)
				}
			}
		})
	}

	tests := []test{
		{
			name:     "successful flow",
			item:     NewItem(context.Background(), 2, nil),
			flow:     DoubleFlow{},
			expected: NewItem(context.Background(), 4, nil),
		},
		{
			name:     "flow with error",
			item:     NewItem(context.Background(), 2, nil),
			flow:     ErrorFlow{},
			expected: NewItem(context.Background(), 0, errors.New("ErrFlowProcessing")),
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}
