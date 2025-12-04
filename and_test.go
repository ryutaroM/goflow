package goflow

import (
	"context"
	"testing"
)

type falseTester struct{}

func (falseTester) Test(ctx context.Context, input int) (bool, error) {
	return false, nil
}

type trueTester struct{}

func (trueTester) Test(ctx context.Context, input int) (bool, error) {
	return true, nil
}

func TestAndPredicator(t *testing.T) {
	type test struct {
		name  string
		left  Predicator[int]
		right Predicator[int]
		input int
		want  bool
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			andPred := AndPredicator(tt.left, tt.right)
			result, err := andPred.Test(context.Background(), tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.want {
				t.Fatalf("got %v, want %v", result, tt.want)
			}
		})
	}

	tests := []test{
		{
			name:  "left true",
			left:  falseTester{},
			right: trueTester{},
			input: 0,
			want:  false,
		},
		{
			name:  "right true",
			left:  trueTester{},
			right: falseTester{},
			input: 0,
			want:  false,
		},
		{
			name:  "both true",
			left:  trueTester{},
			right: trueTester{},
			input: 0,
			want:  true,
		},
		{
			name:  "both false",
			left:  falseTester{},
			right: falseTester{},
			input: 0,
			want:  false,
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}
