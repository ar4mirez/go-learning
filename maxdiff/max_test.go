package maxdiff_test

import (
	"testing"

	"github.com/ar4mirez/go-learning/maxdiff"
)

func TestMaxDiff(t *testing.T) {
	tt := []struct {
		name   string
		input  []int32
		output int32
	}{
		{name: "scenario 1", input: []int32{2, 4, 5, 2, 8}, output: 6},
		{name: "scenario 1", input: []int32{9, 4, 5, 2, 5}, output: 3},
		{name: "scenario 1", input: []int32{10, 14, 15, 1, 1}, output: 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := maxdiff.MaxDiff(tc.input)
			if o != tc.output {
				t.Errorf("expected %d and got %d", tc.output, o)
			}
		})
	}
}
