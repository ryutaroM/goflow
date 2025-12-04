package goflow

import (
	"context"
	"testing"
)

func runTest[T comparable](t *testing.T, value T) {
	item := NewItem(context.Background(), value, nil)
	if item.Value() != value {
		t.Fatal("failed assertion")
	}
}
func TestNewItem(t *testing.T) {
	type test struct {
		name  string
		value any
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			runTest(t, tt.value)
		})
	}

	tests := []test{}
	for _, ts := range tests {
		do(ts, t)
	}
}
