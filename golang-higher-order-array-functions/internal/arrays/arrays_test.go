package arrays

import (
	"fmt"
	"testing"
)

func TestFilterInts(t *testing.T) {
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

func TestFilterObjects(t *testing.T) {
	type Book struct {
		title string
	}

	tests := []struct {
		filter   func(x Book) bool
		input    []Book
		expected []Book
	}{
		{
			func(x Book) bool { return x.title != "book1" },
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
			[]Book{{title: "book2"}, {title: "book3"}},
		},
		{
			func(x Book) bool { return x != Book{title: "book1"} },
			[]Book{{title: "book1"}},
			[]Book{},
		},
		{
			func(x Book) bool { return x.title != "" },
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
		},
		{
			func(x Book) bool { return x.title != "" },
			[]Book{},
			[]Book{},
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
				t.Errorf("Wrong element found. Expected: '%s' found '%s'", result, expected)
			}
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		mutation func(x int32) string
		input    []int32
		expected []string
	}{
		{
			func(x int32) string { return fmt.Sprint(x) },
			[]int32{12345, 123, 125},
			[]string{"12345", "123", "125"},
		},
	}

	for _, test := range tests {
		test := test
		res := Map(test.input, test.mutation)

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%s' found '%s'", result, expected)
			}
		}
	}
}

func TestFirstOrDefault(t *testing.T) {
	tests := []struct {
		inputArr     []int32
		defaultValue int32
		expected     int32
	}{
		{
			[]int32{12345, 123, 125},
			0,
			12345,
		},
		{
			[]int32{},
			0,
			0,
		},
		{
			[]int32{},
			5,
			5,
		},
	}

	for _, test := range tests {
		test := test
		res := FirstOrDefault(test.inputArr, test.defaultValue)

		if res != test.expected {
			t.Errorf("Wrong value returned. Expected: %d found: %d", test.expected, res)
		}
	}
}
