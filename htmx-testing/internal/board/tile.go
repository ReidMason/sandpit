package board

import (
	"fmt"
	"math/rand"
)

const (
	Empty = iota
	Water
	Grass
	Sand
	Forest
)

type TileType int

type Tile struct {
	north         *Tile
	east          *Tile
	south         *Tile
	west          *Tile
	tileType      TileType
	possibilities []TileType
	entropy       int
}

func NewTile() *Tile {
	possibilities := []TileType{Water, Grass, Sand, Forest}
	return &Tile{
		possibilities: possibilities,
		entropy:       len(possibilities),
	}
}

func (t Tile) display() string {
	return fmt.Sprint(t.tileType)
}

func (t *Tile) collapse() {
	idx := rand.Intn(len(t.possibilities))
	t.tileType = t.possibilities[idx]
	t.possibilities = []TileType{t.tileType}
	t.entropy = 0

	possibilities := t.getPossibilities()
	t.north.constrain(possibilities)
	t.east.constrain(possibilities)
	t.south.constrain(possibilities)
	t.west.constrain(possibilities)
}

func (t Tile) getPossibilities() map[TileType]bool {
	possibilities := make(map[TileType]bool)
	for _, p := range t.possibilities {
		possibilities[p] = true
	}

	return possibilities
}

func (t *Tile) constrain(neighbourPossibilities map[TileType]bool) {
	if t == nil || t.entropy <= 0 {
		return
	}

	_, hasWater := neighbourPossibilities[Water]
	_, hasGrass := neighbourPossibilities[Grass]
	_, hasSand := neighbourPossibilities[Sand]
	_, hasForest := neighbourPossibilities[Forest]

	if !hasWater && !hasGrass && !hasSand {
		t.possibilities = t.filterPossibilties(Sand)
	}

	if !hasGrass && !hasForest {
		t.possibilities = t.filterPossibilties(Forest)
	}

	if !hasSand && !hasForest && !hasGrass {
		t.possibilities = t.filterPossibilties(Grass)
	}

	if !hasSand && !hasWater {
		t.possibilities = t.filterPossibilties(Water)
	}

	constrained := len(t.possibilities) != t.entropy
	t.entropy = len(t.possibilities)

	if constrained {
		possibilities := t.getPossibilities()
		t.north.constrain(possibilities)
		t.east.constrain(possibilities)
		t.south.constrain(possibilities)
		t.west.constrain(possibilities)
	}
}

func (t Tile) filterPossibilties(tileType TileType) []TileType {
	newPossibilities := make([]TileType, 0)
	for _, val := range t.possibilities {
		if val != tileType {
			newPossibilities = append(newPossibilities, val)
		}
	}
	return newPossibilities
}
