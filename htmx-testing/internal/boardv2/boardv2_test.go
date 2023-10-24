package boardv2

import (
	"math/rand"
	"testing"
)

// V1 6.2
// V2 5.5
// V3 1.4
func BenchmarkBoard(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	board := New(100, r)
	for board.Iter() {
	}
}
