package main

import (
	"context"
	"fmt"

	"github.com/ryutaroM/goflow"
)

type isGreaterThan5 struct{}

func (p isGreaterThan5) Test(ctx context.Context, input int) (bool, error) {
	return input > 5, nil
}

type isGreaterThan15 struct{}

func (p isGreaterThan15) Test(ctx context.Context, input int) (bool, error) {
	return input > 15, nil
}

type addTenFlow struct{}

func (f addTenFlow) Process(ctx context.Context, input int) (int, error) {
	return input + 10, nil
}

type subtractFiveFlow struct{}

func (f subtractFiveFlow) Process(ctx context.Context, input int) (int, error) {
	return input - 5, nil
}

type AddAFlow struct{}

func (f AddAFlow) Process(ctx context.Context, input string) (string, error) {
	return input + "goflow!", nil
}

func main() {

	item, err := goflow.NewItem(context.Background(), "Hello, ", nil).
		Pipe(AddAFlow{}).Result()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(item)

	b := func(i goflow.Item[int]) goflow.Item[int] {
		return i.Branch(
			goflow.AndPredicator(
				goflow.NotPredicator(isGreaterThan15{}),
				isGreaterThan5{},
			),
			addTenFlow{},
			subtractFiveFlow{},
		)
	}

	result7, err := b(goflow.NewItem(context.Background(), 7, nil)).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	result15, err := b(goflow.NewItem(context.Background(), 15, nil)).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("-----Nested Branching Result-----")
	fmt.Println("Result: ", result7)  // Should print: Result:  17
	fmt.Println("Result: ", result15) // Should print: Result:  25
}
