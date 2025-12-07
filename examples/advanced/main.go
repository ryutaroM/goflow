package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ryutaroM/goflow"
)

type data struct {
	value int
}
type steps struct {
	name    string
	depends []string
	flow    goflow.Flow[data]
}

type dag struct {
	steps   map[string]steps
	results map[string]data
	mu      sync.Mutex
}

func NewDAG() *dag {
	return &dag{
		steps:   make(map[string]steps),
		results: make(map[string]data),
	}
}

func (d *dag) addStep(s steps, deps []string, f goflow.Flow[data]) {
	d.steps[s.name] = steps{
		name:    s.name,
		depends: deps,
		flow:    f,
	}
}

func (d *dag) topologicalSort() ([]string, error) {
	inDegrees := make(map[string]int)
	queue := []string{}
	result := []string{}

	for name := range d.steps {
		inDegrees[name] = 0
	}

	for _, step := range d.steps {
		for range step.depends {
			inDegrees[step.name]++
		}
	}

	for k, v := range inDegrees {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, step := range d.steps {
			for _, dep := range step.depends {
				if dep == current {
					inDegrees[step.name]--
					if inDegrees[step.name] == 0 {
						queue = append(queue, step.name)
					}
				}
			}
		}
	}

	if len(result) != len(d.steps) {
		return nil, fmt.Errorf("cycle detected in DAG")
	}

	return result, nil
}

func (d *dag) execute(ctx context.Context, input data) error {
	ordered, err := d.topologicalSort()
	if err != nil {
		return err
	}

	for _, stepName := range ordered {
		step := d.steps[stepName]
		var in data

		if len(step.depends) == 0 {
			in = input
		} else {
			var merged data
			for _, dep := range step.depends {
				merged.value += d.results[dep].value
			}
			in = merged
		}

		if step.flow != nil {
			output, err := goflow.NewItem(ctx, in, nil).
				Pipe(step.flow).
				Result()
			if err != nil {
				return fmt.Errorf("step %s failed: %w", stepName, err)
			}
			d.mu.Lock()
			d.results[stepName] = output
			d.mu.Unlock()
		}
	}

	return nil
}

type flow1 struct{}

func (f flow1) Process(ctx context.Context, input data) (data, error) {
	return data{value: input.value + 1}, nil
}

type flow2 struct{}

func (f flow2) Process(ctx context.Context, input data) (data, error) {
	return data{value: input.value + 2}, nil
}

type flow3 struct{}

func (f flow3) Process(ctx context.Context, input data) (data, error) {
	return data{value: input.value + 3}, nil
}

type flow4 struct{}

func (f flow4) Process(ctx context.Context, input data) (data, error) {
	return data{value: input.value + 4}, nil
}

type slowFlow struct{}

func (f slowFlow) Process(ctx context.Context, input data) (data, error) {
	select {
	case <-time.After(5 * time.Second):
		return input, nil
	case <-ctx.Done():
		return data{}, ctx.Err()
	}
}

type multiplyByHundredFlow struct{}

func (f multiplyByHundredFlow) Process(ctx context.Context, input data) (data, error) {
	return data{value: input.value * 100}, nil
}

func main() {
	dag := NewDAG()
	dag.addStep(steps{name: "Step4"}, []string{"Step2", "Step3"}, flow4{})
	dag.addStep(steps{name: "Step1"}, []string{}, flow1{})
	dag.addStep(steps{name: "Step2"}, []string{"Step1"}, flow2{})
	dag.addStep(steps{name: "Step3"}, []string{"Step1"}, flow3{})

	err := dag.execute(context.Background(), data{value: 0})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("-----DAG Execution Result-----")
	for name, result := range dag.results {
		fmt.Printf("%s: %d\n", name, result.value)
	}

	withTimeoutCtx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	result1, err := goflow.NewItem(withTimeoutCtx, data{value: 1}, nil).
		Pipe(slowFlow{}).
		Pipe(multiplyByHundredFlow{}).
		Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result without timeout:", result1)
}
