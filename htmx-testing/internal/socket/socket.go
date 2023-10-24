package socket

type Socket int

const (
	Grass = iota
	Forest
	Water
	Sand

	WaterT
	SandT
	GrassT

	WaterSandW
	WaterSandE

	WaterSandCornerN
	WaterSandCornerW

	SandWaterCornerN
	SandWaterCornerW

	SandGrassW
	SandGrassE
)

var SocketConstraints = map[Socket]map[Socket]bool{
	Grass:  {Grass: true, Forest: true, Sand: true},
	Forest: {Forest: true, Grass: true},
	Water:  {Water: true, WaterT: true},
	Sand:   {Sand: true, SandT: true, Grass: true},

	WaterSandW: {WaterSandE: true},

	WaterSandCornerN: {WaterSandW: true},
	WaterSandCornerW: {WaterSandE: true},

	SandWaterCornerN: {WaterSandE: true},
	SandWaterCornerW: {WaterSandW: true},
}

func CanConnect(socket1, socket2 Socket) bool {
	compatibleSockets, found := SocketConstraints[socket1]
	if found {
		_, found = compatibleSockets[socket2]
		if found {
			return true
		}
	}

	compatibleSockets, found = SocketConstraints[socket2]
	if found {
		_, found = compatibleSockets[socket1]
		if found {
			return true
		}
	}

	return false
}
