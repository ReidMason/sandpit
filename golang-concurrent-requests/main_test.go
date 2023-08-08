package main

import "testing"

func BenchmarkMain(b *testing.B) {
	main()
}

func TestMain(t *testing.T) {
	main()
}
