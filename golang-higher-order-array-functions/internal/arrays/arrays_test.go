package arrays

import (
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		filter   func(x int32) bool
		input    []int32
		expected []int32
	}{
		{
			func(x int32) bool { return x != 12345 },
			[]int32{12345, 123, 125},
			[]int32{123, 125},
		},
		{
			func(x int32) bool { return x != 12345 },
			[]int32{12345},
			[]int32{},
		},
		{
			func(x int32) bool { return x != 0 },
			[]int32{12345, 123, 125},
			[]int32{12345, 123, 125},
		},
		{
			func(x int32) bool { return x != 0 },
			[]int32{},
			[]int32{},
		},
	}

	for _, test := range tests {
		test := test
		res := Filter(test.input, test.filter)

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%d' found '%d'", result, expected)
			}
		}
	}
}
