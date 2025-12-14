package dag

import (
	"reflect"
	"testing"

	"github.com/ryutaroM/goflow"
)

type exStep struct {
	name    string
	depends []string
	flow    goflow.Flow[int]
}

func (s *exStep) Name() string {
	return s.name
}

func (s *exStep) Depends() []string {
	return s.depends
}

func (s *exStep) Flow() goflow.Flow[int] {
	return s.flow
}

func newExStep(name string, depends []string, flow goflow.Flow[int]) *exStep {
	return &exStep{
		name:    name,
		depends: depends,
		flow:    flow,
	}
}

func TestDAGTopologicalSort(t *testing.T) {
	type test struct {
		name    string
		dag     *DAG[int]
		wantErr bool
		want    []string
	}

	do := func(tt test, t *testing.T) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dag.TopologicalSort()
			if (err != nil) != tt.wantErr {
				t.Fatalf("got error %v, wantErr %v", err, tt.wantErr)
			}

			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}

	tests := []test{
		{
			name: "simple DAG",
			dag: NewDAG(
				func() map[string]Step[int] {
					steps := make(map[string]Step[int])
					steps["A"] = newExStep("A", []string{}, nil)
					steps["B"] = newExStep("B", []string{"A"}, nil)
					steps["C"] = newExStep("C", []string{"A"}, nil)
					steps["D"] = newExStep("D", []string{"B", "C"}, nil)
					return steps
				}(),
			),
			wantErr: false,
			want:    []string{"A", "B", "C", "D"},
		},
	}

	for _, tt := range tests {
		do(tt, t)
	}
}
