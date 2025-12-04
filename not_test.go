package goflow

import (
	"context"
	"testing"
)

func TestNotPredicator(t *testing.T) {
	type test struct {
		name     string
		inner    Predicator[int]
		input    int
		expected bool
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			notPred := NotPredicator(tt.inner)
			result, err := notPred.Test(context.Background(), tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
		})
	}

	tests := []test{
		{
			name:     "inner true",
			inner:    trueTester{},
			input:    0,
			expected: false,
		},
		{
			name:     "inner false",
			inner:    falseTester{},
			input:    0,
			expected: true,
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}
