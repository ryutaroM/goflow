package main

import (
	"context"
	"fmt"

	"github.com/ryutaroM/goflow"
)

type trueTester struct{}

func (p trueTester) Test(ctx context.Context, input int) (bool, error) {
	return true, nil
}

type doubleFlow struct{}

func (f doubleFlow) Process(ctx context.Context, input int) (int, error) {
	return input * 2, nil
}

type addTenFlow struct{}

func (f addTenFlow) Process(ctx context.Context, input int) (int, error) {
	return input + 10, nil
}

func main() {
	pred := goflow.NotPredicator(
		trueTester{},
	)

	result, err := goflow.NewItem(context.Background(), 4, nil).
		Branch(
			pred,
			doubleFlow{},
			addTenFlow{},
		).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Result: ", result) // Should print: Result:  14
}
