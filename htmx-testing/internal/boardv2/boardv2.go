package boardv2

import (
	"htmx-testing/internal/cell"
	"htmx-testing/internal/tile"
	"math/rand"
)

type Board struct {
	r     *rand.Rand
	cells [][]*cell.Cell
	size  int
}

func New(size int, r *rand.Rand) *Board {
	allTiles := tile.GetAllTiles()
	cells := make([][]*cell.Cell, 0, size)
	for i := 0; i < size; i++ {
		childArr := make([]*cell.Cell, 0, size)
		for i := 0; i < size; i++ {
			childArr = append(childArr, cell.New(r, allTiles))
		}
		cells = append(cells, childArr)
	}

	// Make each cell aware of it's neighbours
	for y, row := range cells {
		for x, cell := range row {
			if y > 0 {
				cell.North = cells[y-1][x]
			}

			if x < size-1 {
				cell.East = cells[y][x+1]
			}

			if y < size-1 {
				cell.South = cells[y+1][x]
			}

			if x > 0 {
				cell.West = cells[y][x-1]
			}
		}
	}

	return &Board{cells: cells, size: size, r: r}
}

func (b Board) Display() [][]TileDisplay {
	display := make([][]TileDisplay, 0, b.size)
	for _, row := range b.cells {
		displayRow := make([]TileDisplay, 0, b.size)
		for _, cell := range row {
			displayRow = append(displayRow, TileDisplay{
				Style: cell.Tile.Style,
				// Content: fmt.Sprint(cell.Entropy()),
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
	lowestCells := make([]*cell.Cell, 0)
	for _, row := range b.cells {
		for _, boardCell := range row {
			if boardCell.Collapsed() || boardCell.Entropy() == 0 {
				continue
			}

			if len(lowestCells) == 0 || lowestCells[0].Entropy() > boardCell.Entropy() {
				lowestCells = []*cell.Cell{boardCell}
			} else if lowestCells[0].Entropy() == boardCell.Entropy() {
				lowestCells = append(lowestCells, boardCell)
			}
		}
	}

	if len(lowestCells) == 0 {
		return false
	}

	idx := b.r.Intn(len(lowestCells))
	randomTile := lowestCells[idx]
	randomTile.Collapse()

	return true
}
