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
const VerticalGradient = ""

func getStyle(tileType TileType) string {
	switch tileType {
	case Water:
		return "bg-blue-400"
	case WaterSandTL:
		return fmt.Sprintf("bg-gradient-to-tl from-red-400 %s to-purple-200", DiagonalGradient)
	case WaterSandTR:
		return fmt.Sprintf("bg-gradient-to-tr from-red-400 %s to-purple-200", DiagonalGradient)
	case WaterSandBR:
		return fmt.Sprintf("bg-gradient-to-br from-red-400 %s to-purple-200", DiagonalGradient)
	case WaterSandBL:
		return fmt.Sprintf("bg-gradient-to-bl from-red-400 %s to-purple-200", DiagonalGradient)
	case WaterSandN:
		return fmt.Sprintf("bg-gradient-to-t from-red-400 %s to-purple-200", VerticalGradient)
	case WaterSandE:
		return fmt.Sprintf("bg-gradient-to-r from-red-400 %s to-purple-200", VerticalGradient)
	case WaterSandS:
		return fmt.Sprintf("bg-gradient-to-b from-red-400 %s to-purple-200", VerticalGradient)
	case WaterSandW:
		return fmt.Sprintf("bg-gradient-to-l from-red-400 %s to-purple-200", VerticalGradient)
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
	north                  *Tile
	east                   *Tile
	south                  *Tile
	west                   *Tile
	tileType               TileType
	r                      rand.Rand
	neighbourPossibilities map[Direction][]TileType
	possibilities          []TileType
	entropy                int
}

func NewTile(r rand.Rand) *Tile {
	possibilities := []TileType{
		Water,
		WaterSandN,
		WaterSandS,
		Grass,
		Sand,
		Forest,
	}

	neighbourPossibilities := map[Direction][]TileType{
		North: append([]TileType{}, possibilities...),
		East:  append([]TileType{}, possibilities...),
		South: append([]TileType{}, possibilities...),
		West:  append([]TileType{}, possibilities...),
	}
	return &Tile{
		neighbourPossibilities: neighbourPossibilities,
		possibilities:          possibilities,
		entropy:                len(possibilities),
		r:                      r,
	}
}

func (t Tile) display() string {
	return fmt.Sprint(t.tileType)
}

func (t *Tile) collapse() {
	idx := t.r.Intn(len(t.possibilities))
	t.tileType = t.possibilities[idx]

	t.possibilities = []TileType{t.tileType}
	t.neighbourPossibilities = TileConstraints[t.tileType]

	t.constrain(t.north, North)
	t.constrain(t.east, East)
	t.constrain(t.south, South)
	t.constrain(t.west, West)

	t.entropy = 0
}

func (t *Tile) constrain(neighbour *Tile, direction Direction) {
	if t == nil || neighbour == nil || t.entropy <= 0 {
		return
	}

	neighbourPossibilities := neighbour.possibilities
	if len(neighbourPossibilities) == 0 {
		return
	}

	ourPossibilities := t.neighbourPossibilities[direction]

	constrained := false
	for _, neighbourPossibility := range neighbourPossibilities {
		contains := false
		for _, possibility := range ourPossibilities {
			if neighbourPossibility == possibility {
				contains = true
				break
			}
		}

		if !contains {
			constrained = true
			neighbour.possibilities = neighbour.filterPossibilties(neighbourPossibility)
			neighbour.entropy = len(neighbour.possibilities)
		}
	}

	if constrained {
		neighbour.constrain(neighbour.north, North)
		neighbour.constrain(neighbour.east, East)
		neighbour.constrain(neighbour.south, South)
		neighbour.constrain(neighbour.west, West)
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
