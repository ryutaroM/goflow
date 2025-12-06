package main

import (
	"context"
	"fmt"

	"github.com/ryutaroM/goflow"
)

type IsEven struct{}

func (p IsEven) Test(ctx context.Context, input int) (bool, error) {
	return input%2 == 0, nil
}

type IsPositive struct{}

func (p IsPositive) Test(ctx context.Context, input int) (bool, error) {
	return input > 0, nil
}

type DoubleFlow struct{}

func (f DoubleFlow) Process(ctx context.Context, input int) (int, error) {
	return input * 2, nil
}

type DivideFlow struct{}

func (f DivideFlow) Process(ctx context.Context, input int) (int, error) {
	return input / 2, nil
}

func main() {
	pred := goflow.AndPredicator(
		IsEven{},
		IsPositive{},
	)

	fmt.Println(pred.Test(context.Background(), 1))
	fmt.Println(pred.Test(context.Background(), 2))
	fmt.Println(pred.Test(context.Background(), -2))

	result, err := goflow.NewItem(context.Background(), 4, nil).
		Branch(pred, DoubleFlow{}, DivideFlow{}).
		Result()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
