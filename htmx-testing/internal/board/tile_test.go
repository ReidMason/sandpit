package board

import "testing"

func TestSomething(t *testing.T) {
	result := findPossibleConnectors(WaterSandN, East)

	for _, r := range result {
		t.Log(r)
	}
}
