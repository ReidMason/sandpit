package socket

type Socket int

const (
	Grass = iota
	Forest
	Water
	Sand
)

var SocketConstraints = map[Socket]map[Socket]bool{
	Grass:  {Grass: true, Forest: true, Sand: true},
	Forest: {Forest: true, Grass: true},
	Water:  {Water: true, Sand: true},
	Sand:   {Sand: true, Water: true, Grass: true},
}

func CanConnect(socket1, socket2 Socket) bool {
	compatibleSockets, found := SocketConstraints[socket1]
	if !found {
		return false
	}

	_, found = compatibleSockets[socket2]
	return found
}
