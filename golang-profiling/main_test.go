package main

import (
	"testing"
)

var items = 2_000
var iterations = 1_000
var targetId = items / 2

type Data struct {
	id   int
	text string
}

func BenchmarkAccessArray(b *testing.B) {
	myArr := []Data{}
	for i := 0; i < items; i++ {
		myArr = append(myArr, Data{id: i, text: "Test"})
	}
	b.ResetTimer()

	for i := 0; i < iterations; i++ {
		for _, x := range myArr {
			if x.id == targetId {
				break
			}
		}
	}
}

func BenchmarkHashMap(b *testing.B) {
	myMap := make(map[int]string)
	for i := 0; i < items; i++ {
		myMap[i] = "Test"
	}
	b.ResetTimer()

	for i := 0; i < iterations; i++ {
		_ = myMap[targetId]
	}

}
