package main

import (
	"context"
	"fmt"
	"ryutaroM/goflow"
	"strconv"
)

type AddAFlow struct{}

func (f AddAFlow) Process(ctx context.Context, input string) (string, error) {
	return input + "goflow!", nil
}

type StringToInt struct{}

func (t StringToInt) Transform(ctx context.Context, input string) (int, error) {
	return strconv.Atoi(input)
}

func main() {

	item, err := goflow.NewItem(context.Background(), "Hello, ", nil).
		Pipe(AddAFlow{}).Result()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(item)
}
