package goflow

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestNewItem(t *testing.T) {
	type test struct {
		name  string
		value any
		want  Item[any]
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			item := NewItem(context.Background(), tt.value, nil)
			if item.err != tt.want.err {
				t.Fatalf("error mismatch: got %v, want %v", item.err, tt.want.err)
			}
			itemCopy := item.Value()
			wantCopy := tt.want.Value()
			if !reflect.DeepEqual(itemCopy, wantCopy) {
				t.Fatalf("got %v, want %v", item, tt.want)
			}
			itemCtx := item.ctx
			wantCtx := tt.want.ctx
			if itemCtx != wantCtx {
				t.Fatalf("context mismatch: got %v, want %v", itemCtx, wantCtx)
			}
		})
	}

	tests := []test{
		{
			name:  "int value",
			value: 42,
			want:  NewItem[any](context.Background(), 42, nil),
		},
		{
			name:  "string value",
			value: "hello",
			want:  NewItem[any](context.Background(), "hello", nil),
		},
		{
			name:  "nil value",
			value: nil,
			want:  NewItem[any](context.Background(), nil, nil),
		},
		{
			name:  "struct value",
			value: struct{ A int }{A: 10},
			want:  NewItem[any](context.Background(), struct{ A int }{A: 10}, nil),
		},
		{
			name:  "slice value",
			value: []int{1, 2, 3},
			want:  NewItem[any](context.Background(), []int{1, 2, 3}, nil),
		},
		{
			name:  "map value",
			value: map[string]int{"a": 1, "b": 2},
			want:  NewItem[any](context.Background(), map[string]int{"a": 1, "b": 2}, nil),
		},
		{
			name:  "pointer value",
			value: &struct{ B string }{B: "test"},
			want:  NewItem[any](context.Background(), &struct{ B string }{B: "test"}, nil),
		},
		{
			name:  "float value",
			value: 3.14,
			want:  NewItem[any](context.Background(), 3.14, nil),
		},
		{
			name:  "boolean value",
			value: true,
			want:  NewItem[any](context.Background(), true, nil),
		},
		{
			name:  "rune value",
			value: 'g',
			want:  NewItem[any](context.Background(), 'g', nil),
		},
	}
	for _, ts := range tests {
		do(ts, t)
	}
}

func TestItemValue(t *testing.T) {
	type test struct {
		name  string
		item  Item[any]
		value any
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.item.Value(), tt.value) {
				t.Fatalf("got %v, want %v", tt.item.Value(), tt.value)
			}
		})
	}

	tests := []test{
		{
			name:  "int value",
			item:  NewItem[any](context.Background(), 42, nil),
			value: 42,
		},
		{
			name:  "string value",
			item:  NewItem[any](context.Background(), "hello", nil),
			value: "hello",
		},
		{
			name:  "nil value",
			item:  NewItem[any](context.Background(), nil, nil),
			value: nil,
		},
		{
			name:  "struct value",
			item:  NewItem[any](context.Background(), struct{ A int }{A: 10}, nil),
			value: struct{ A int }{A: 10},
		},
		{
			name:  "slice value",
			item:  NewItem[any](context.Background(), []int{1, 2, 3}, nil),
			value: []int{1, 2, 3},
		},
		{
			name:  "map value",
			item:  NewItem[any](context.Background(), map[string]int{"a": 1, "b": 2}, nil),
			value: map[string]int{"a": 1, "b": 2},
		},
		{
			name:  "pointer value",
			item:  NewItem[any](context.Background(), &struct{ B string }{B: "test"}, nil),
			value: &struct{ B string }{B: "test"},
		},
		{
			name:  "float value",
			item:  NewItem[any](context.Background(), 3.14, nil),
			value: 3.14,
		},
		{
			name:  "boolean value",
			item:  NewItem[any](context.Background(), true, nil),
			value: true,
		},
		{
			name:  "rune value",
			item:  NewItem[any](context.Background(), 'g', nil),
			value: 'g',
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}

func TestItemError(t *testing.T) {
	type test struct {
		name  string
		item  Item[any]
		error error
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			if tt.item.Error() != tt.error {
				t.Fatalf("got %v, want %v", tt.item.Error(), tt.error)
			}
		})
	}

	tests := []test{
		{
			name:  "no error",
			item:  NewItem[any](context.Background(), 42, nil),
			error: nil,
		},
		{
			name:  "with error",
			item:  NewItem[any](context.Background(), 42, context.Canceled),
			error: context.Canceled,
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}

func TestItemResult(t *testing.T) {
	type test struct {
		name     string
		item     Item[any]
		value    any
		errorVal error
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			v, err := tt.item.Result()
			if v != tt.value {
				t.Fatalf("got value %v, want %v", v, tt.value)
			}
			if err != tt.errorVal {
				t.Fatalf("got error %v, want %v", err, tt.errorVal)
			}
		})
	}

	tests := []test{
		{
			name:     "no error",
			item:     NewItem[any](context.Background(), 42, nil),
			value:    42,
			errorVal: nil,
		},
		{
			name:     "with error",
			item:     NewItem[any](context.Background(), 42, context.Canceled),
			value:    42,
			errorVal: context.Canceled,
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}

type simpleFlow struct{}

func (sf simpleFlow) Process(ctx context.Context, input any) (any, error) {
	return input.(int) * 2, nil
}

func TestItemPipe(t *testing.T) {
	type test struct {
		name string
		item Item[any]
		flow Flow[any]
		want Item[any]
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.item.Pipe(tt.flow)
			gotValue := got.Value()
			wantValue := tt.want.Value()
			if !reflect.DeepEqual(gotValue, wantValue) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}

	tests := []test{
		{
			name: "simple flow",
			item: NewItem[any](context.Background(), 5, nil),
			flow: simpleFlow{},
			want: NewItem[any](context.Background(), 10, nil),
		},
		{
			name: "flow with error",
			item: NewItem[any](context.Background(), 5, context.Canceled),
			flow: simpleFlow{},
			want: NewItem[any](context.Background(), 5, context.Canceled),
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}

type simplePredicator struct{}

func (sp simplePredicator) Test(ctx context.Context, input any) (bool, error) {
	return true, nil
}

type simplePredicatorFalse struct{}

func (sp simplePredicatorFalse) Test(ctx context.Context, input any) (bool, error) {
	return false, nil
}

type leftFlow struct{}

func (lf leftFlow) Process(ctx context.Context, input any) (any, error) {
	return "left", nil
}

type rightFlow struct{}

func (rf rightFlow) Process(ctx context.Context, input any) (any, error) {
	return "right", nil
}

func TestBranch(t *testing.T) {
	type test struct {
		name  string
		pd    Predicator[any]
		left  Flow[any]
		right Flow[any]
		item  Item[any]
		want  Item[any]
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.item.Branch(tt.pd, tt.left, tt.right)
			gotValue := got.Value()
			wantValue := tt.want.Value()
			if !reflect.DeepEqual(gotValue, wantValue) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}

	tests := []test{
		{
			name:  "branch true",
			pd:    simplePredicator{},
			left:  leftFlow{},
			right: rightFlow{},
			item:  NewItem[any](context.Background(), "test", nil),
			want:  NewItem[any](context.Background(), "left", nil),
		},
		{
			name:  "branch false",
			pd:    simplePredicatorFalse{},
			left:  leftFlow{},
			right: rightFlow{},
			item:  NewItem[any](context.Background(), "test", nil),
			want:  NewItem[any](context.Background(), "right", nil),
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}

type IntToStringTransformer struct{}

func (t IntToStringTransformer) Transform(ctx context.Context, input int) (string, error) {
	return fmt.Sprintf("%d", input), nil
}

func TestTransform(t *testing.T) {
	type test struct {
		name string
		item Item[int]
		want Item[string]
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			transformer := IntToStringTransformer{}
			got := Transform(tt.item, transformer)
			gotValue := got.Value()
			wantValue := tt.want.Value()
			if !reflect.DeepEqual(gotValue, wantValue) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}

	tests := []test{
		{
			name: "transform int to string",
			item: NewItem[int](context.Background(), 123, nil),
			want: NewItem[string](context.Background(), "123", nil),
		},
	}

	for _, ts := range tests {
		do(ts, t)
	}
}
