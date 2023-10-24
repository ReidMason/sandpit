package tile

import "htmx-testing/internal/socket"

type Tile struct {
	Style   string
	Sockets [4]socket.Socket
}

var Blank = Tile{
	Style: "bg-slate-100",
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

	return []Tile{grass, forest, sand, water}
}
