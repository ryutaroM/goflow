package main

import (
	"context"
	"fmt"

	"github.com/ryutaroM/goflow"
)

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
}
