package board

import (
	"math/rand"
	"testing"
)

func BenchmarkBoard(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	board := New(100, *r)
	for board.Iter() {
	}
}
