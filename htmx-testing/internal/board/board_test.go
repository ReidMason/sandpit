package board

import (
	"math/rand"
	"testing"
)

// V1 6.2
// V2 5.5
func BenchmarkBoard(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	board := New(100, *r)
	for board.Iter() {
	}
}
