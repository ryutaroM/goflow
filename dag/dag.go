package dag

import (
	"github.com/ryutaroM/goflow"
)

type Step[T any] interface {
	Name() string
	Depends() []string
	Flow() goflow.Flow[T]
}

type Data[T any] struct {
	Values map[string]T
}

type DAG[T any] struct {
	steps map[string]Step[T]
}

func NewDAG[T any](s map[string]Step[T]) *DAG[T] {
	return &DAG[T]{
		steps: s,
	}
}

func (d *DAG[T]) AddStep(s Step[T]) {
	d.steps[s.Name()] = s
}

func (d *DAG[T]) TopologicalSort() ([]string, error) {
	inDegrees := make(map[string]int)
	queue := []string{}
	result := []string{}

	for name := range d.steps {
		inDegrees[name] = 0
	}

	for _, step := range d.steps {
		for range step.Depends() {
			inDegrees[step.Name()]++
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

		for _, dep := range d.steps[current].Depends() {
			inDegrees[dep]--
			if inDegrees[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	if len(result) != len(d.steps) {
		return nil, wrapError(ErrorCycleDetectedNum)
	}

	return result, nil
}
