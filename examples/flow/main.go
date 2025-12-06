package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ryutaroM/goflow"
)

type TrimFlow struct{}

func (f TrimFlow) Process(ctx context.Context, input string) (string, error) {
	return strings.TrimSpace(input), nil
}

type UpperFlow struct{}

func (f UpperFlow) Process(ctx context.Context, input string) (string, error) {
	return strings.ToUpper(input), nil
}

type AddPrefixFlow struct{}

func (f AddPrefixFlow) Process(ctx context.Context, input string) (string, error) {
	return "PREFIX_" + input, nil
}

func main() {
	result, err := goflow.NewItem(context.Background(), "   hello goflow   ", nil).
		Pipe(TrimFlow{}).
		Pipe(UpperFlow{}).
		Pipe(AddPrefixFlow{}).
		Result()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Result: ", result) // Should print: Result:  PREFIX_HELLO GOFLOW
}
