package goflow

import (
	"context"
	"testing"
)

func TestOrPredicator(t *testing.T) {
	type test struct {
		name     string
		left     Predicator[int]
		right    Predicator[int]
		input    int
		expected bool
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			orPred := OrPredicator(tt.left, tt.right)
			result, err := orPred.Test(context.Background(), tt.input)
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
			name:     "left true",
			left:     trueTester{},
			right:    falseTester{},
			input:    0,
			expected: true,
		},
		{
			name:     "right true",
			left:     falseTester{},
			right:    trueTester{},
			input:    0,
			expected: true,
		},
		{
			name:     "both false",
			left:     falseTester{},
			right:    falseTester{},
			input:    0,
			expected: false,
		},
		{
			name:     "both true",
			left:     trueTester{},
			right:    trueTester{},
			input:    0,
			expected: true,
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}
