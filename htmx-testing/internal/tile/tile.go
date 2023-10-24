package tile

import (
	"htmx-testing/internal/socket"
	"strings"
)

type Tile struct {
	Style   string
	Sockets [4]socket.Socket
}

var Blank = Tile{
	Style: "bg-slate-500",
}

func GetAllTiles() []Tile {
	grass := Tile{
		Sockets: [4]socket.Socket{
			socket.Grass,
			socket.Grass,
			socket.Grass,
			socket.Grass,
		},
		Style: "bg-green-200",
	}

	forest := Tile{
		Sockets: [4]socket.Socket{
			socket.Forest,
			socket.Forest,
			socket.Forest,
			socket.Forest,
		},
		Style: "bg-green-400",
	}

	sand := Tile{
		Sockets: [4]socket.Socket{
			socket.Sand,
			socket.Sand,
			socket.Sand,
			socket.Sand,
		},
		Style: "bg-orange-200",
	}

	water := Tile{
		Sockets: [4]socket.Socket{
			socket.Water,
			socket.Water,
			socket.Water,
			socket.Water,
		},
		Style: "bg-blue-400",
	}

	waterSand := Tile{
		Sockets: [4]socket.Socket{
			socket.WaterT,
			socket.WaterSandE,
			socket.SandT,
			socket.WaterSandW,
		},
		Style: "bg-gradient-to-t from-orange-200 to-blue-400",
	}

	waterCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.WaterSandCornerN,
			socket.SandT,
			socket.SandT,
			socket.WaterSandCornerW,
		},
		Style: "bg-gradient-to-tl from-orange-200 from-50% to-blue-400",
	}

	sandCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.SandWaterCornerN,
			socket.WaterT,
			socket.WaterT,
			socket.SandWaterCornerW,
		},
		Style: "bg-gradient-to-tl from-blue-400 from-50% to-orange-200",
	}

	// sandGrass := Tile{
	// 	Sockets: [4]socket.Socket{
	// 		socket.SandT,
	// 		socket.SandGrassE,
	// 		socket.GrassT,
	// 		socket.SandGrassW,
	// 	},
	// 	Style: "bg-gradient-to-t from-green-200 to-orange-200",
	// }

	// sandGrassCorner := Tile{
	// 	Sockets: [4]socket.Socket{
	// 		socket.SandGrassW,
	// 		socket.GrassT,
	// 		socket.GrassT,
	// 		socket.SandGrassE,
	// 	},
	// 	Style: "bg-gradient-to-tl from-green-200 from-50% to-orange-400",
	// }

	// sandCorner := Tile{
	// 	Sockets: [4]socket.Socket{
	// 		socket.WaterSandDL,
	// 		socket.WaterT,
	// 		socket.WaterT,
	// 		socket.WaterSandDR,
	// 	},
	// 	Style: "bg-gradient-to-tl from-blue-400 from-50% to-orange-200",
	// }

	tiles := []Tile{grass, forest, sand, water}

	rotatableTiles := []Tile{waterSand, waterCorner, sandCorner} //, sandGrass} //, sandGrassCorner}
	for _, tile := range rotatableTiles {
		for i := 0; i < 4; i++ {
			tiles = append(tiles, rotate(tile, i))
		}
	}

	return tiles
}

func rotate(tile Tile, rotations int) Tile {
	totalSockets := len(tile.Sockets)
	newSockets := [4]socket.Socket{}
	for i := 0; i < totalSockets; i++ {
		newSockets[i] = tile.Sockets[(i-rotations+totalSockets)%totalSockets]
	}
	tile.Sockets = newSockets
	tile.Style = getRotateClass(rotations, tile.Style)
	return tile
}

func getRotateClass(rotations int, style string) string {
	if strings.Contains(style, "bg-gradient-to-tl") {
		return rotateDiagonalGradientClass(rotations, style)
	} else if strings.Contains(style, "bg-gradient-to-t") {
		return rotateTopGradientClass(rotations, style)
	}

	return style
}

func rotateDiagonalGradientClass(rotations int, style string) string {
	rotationClasses := []string{"bg-gradient-to-tl", "bg-gradient-to-bl", "bg-gradient-to-br", "bg-gradient-to-tr"}
	count := len(rotationClasses)
	newRotation := rotationClasses[(0-rotations+count)%count]

	return strings.Replace(style, rotationClasses[0], newRotation, 1)
}

func rotateTopGradientClass(rotations int, style string) string {
	rotationClasses := []string{"bg-gradient-to-t", "bg-gradient-to-l", "bg-gradient-to-b", "bg-gradient-to-r"}
	count := len(rotationClasses)
	newRotation := rotationClasses[(0-rotations+count)%count]

	return strings.Replace(style, rotationClasses[0], newRotation, 1)
}
