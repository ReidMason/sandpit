package board

import (
	"fmt"
	"math/rand"
)

const (
	North = iota
	South
	East
	West
)

type Direction int

const (
	Empty       = "Empty"
	Water       = "Water"
	WaterSandTL = "Water to Sand Top Left"
	WaterSandTR = "Water to Sand Top Right"
	WaterSandBR = "Water to Sand Bottom Right"
	WaterSandBL = "Water to Sand Bottom Left"
	WaterSandN  = "Water to Sand North"
	WaterSandE  = "Water to Sand East"
	WaterSandS  = "Water to Sand South"
	WaterSandW  = "Water to Sand West"
	Grass       = "Grass"
	Sand        = "Sand"
	Forest      = "Forest"
)

const DiagonalGradient = "from-40% to-50%"
const VerticalGradient = "from-80%"

func getStyle(tileType TileType) string {
	switch tileType {
	case Water:
		return "bg-blue-400"
	case WaterSandTL:
		return fmt.Sprintf("bg-gradient-to-tl from-blue-400 %s to-orange-200", DiagonalGradient)
	case WaterSandTR:
		return fmt.Sprintf("bg-gradient-to-tr from-blue-400 %s to-orange-200", DiagonalGradient)
	case WaterSandBR:
		return fmt.Sprintf("bg-gradient-to-br from-blue-400 %s to-orange-200", DiagonalGradient)
	case WaterSandBL:
		return fmt.Sprintf("bg-gradient-to-bl from-blue-400 %s to-orange-200", DiagonalGradient)
	case WaterSandN:
		return fmt.Sprintf("bg-gradient-to-t from-blue-400 %s to-orange-200", VerticalGradient)
	case WaterSandE:
		return fmt.Sprintf("bg-gradient-to-r from-blue-400 %s to-orange-200", VerticalGradient)
	case WaterSandS:
		return fmt.Sprintf("bg-gradient-to-b from-blue-400 %s to-orange-200", VerticalGradient)
	case WaterSandW:
		return fmt.Sprintf("bg-gradient-to-l from-blue-400 %s to-orange-200", VerticalGradient)
	case Grass:
		return "bg-green-200"
	case Sand:
		return "bg-orange-200"
	case Forest:
		return "bg-green-500"
	default:
		return "bg-slate-100"
	}
}

type TileType string

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
	possibilities := []TileType{
		Water,
		// WaterSandTL,
		// WaterSandTR,
		// WaterSandBR,
		// WaterSandBL,
		// WaterSandN,
		// WaterSandS,
		Grass,
		Sand,
		Forest,
	}
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

	t.north.constrain(t, North)
	t.east.constrain(t, East)
	t.south.constrain(t, South)
	t.west.constrain(t, West)
}

func findPossibleConnectors(tileType TileType, direction Direction) []TileType {
	foundConnectors := make([]TileType, 0)
	for _, constraint := range TileConstraints {
		if constraint.tileType != tileType {
			continue
		}

		directionType := constraint.north
		if direction == North {
			directionType = constraint.north
		} else if direction == South {
			directionType = constraint.south
		} else if direction == East {
			directionType = constraint.east
		} else if direction == West {
			directionType = constraint.west
		}

		foundConnectors = append(foundConnectors, directionType...)
	}

	return foundConnectors
}

func (t *Tile) constrain(neighbour *Tile, direction Direction) {
	if t == nil || t.entropy <= 0 {
		return
	}

	neighbourPossibilities := neighbour.possibilities
	connectors := make(map[TileType]bool)
	for _, possibility := range neighbourPossibilities {
		possibleConnectors := findPossibleConnectors(possibility, direction)
		for _, connector := range possibleConnectors {
			connectors[connector] = true
		}
	}

	newPossibilities := make([]TileType, 0)
	for _, possibility := range t.possibilities {
		_, found := connectors[possibility]
		if found {
			newPossibilities = append(newPossibilities, possibility)
		}
	}

	constrained := len(newPossibilities) != t.entropy
	t.entropy = len(newPossibilities)
	t.possibilities = newPossibilities

	if constrained {
		t.north.constrain(t, North)
		t.east.constrain(t, East)
		t.south.constrain(t, South)
		t.west.constrain(t, West)
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
