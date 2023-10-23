package board

import (
	"fmt"
	"math/rand"
)

type Board struct {
	r     rand.Rand
	tiles [][]*Tile
	size  int
}

func New(size int, r rand.Rand) *Board {
	// Init the tiles
	tiles := make([][]*Tile, 0, size)
	for i := 0; i < size; i++ {
		childArr := make([]*Tile, 0, size)
		for i := 0; i < size; i++ {
			childArr = append(childArr, NewTile(r))
		}
		tiles = append(tiles, childArr)
	}

	// Make each child aware of it's neighbour
	for y, row := range tiles {
		for x, tile := range row {
			if y > 0 {
				tile.north = tiles[y-1][x]
			}

			if x < size-1 {
				tile.east = tiles[y][x+1]
			}

			if y < size-1 {
				tile.south = tiles[y+1][x]
			}

			if x > 0 {
				tile.west = tiles[y][x-1]
			}
		}
	}

	return &Board{tiles: tiles, size: size, r: r}
}

func (b Board) Display() [][]TileDisplay {
	display := make([][]TileDisplay, 0, b.size)
	for _, row := range b.tiles {
		displayRow := make([]TileDisplay, 0, b.size)
		for _, tile := range row {
			displayRow = append(displayRow, TileDisplay{
				Style:   getStyle(tile.tileType),
				Content: fmt.Sprint(tile.entropy),
				// Content: strings.Join(Map(tile.possibilities,
				// func(x TileType) string { return fmt.Sprint(x) }), "-"),
			})
		}
		display = append(display, displayRow)
	}

	return display
}

type TileDisplay struct {
	Content string
	Style   string
}

func (b *Board) Iter() bool {
	lowestTiles := make([]*Tile, 0)
	lowestEntropy := 0
	for _, row := range b.tiles {
		for _, tile := range row {
			if tile.entropy == 0 {
				continue
			}

			if lowestEntropy == 0 || tile.entropy < lowestEntropy {
				lowestTiles = []*Tile{tile}
				lowestEntropy = tile.entropy
			} else if lowestEntropy == tile.entropy {
				lowestTiles = append(lowestTiles, tile)
			}
		}
	}

	if len(lowestTiles) == 0 {
		return false
	}

	idx := b.r.Intn(len(lowestTiles))
	randomTile := lowestTiles[idx]
	randomTile.collapse()

	return true
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
