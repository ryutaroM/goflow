package main

import (
	"context"
	"strconv"
)

type StringToInt struct{}

func (t StringToInt) Transform(ctx context.Context, input string) (int, error) {
	return strconv.Atoi(input)
}

func main() {
	transformer := StringToInt{}
	result, err := transformer.Transform(context.Background(), "12345")
	if err != nil {
		panic(err)
	}
	println(result)
}
