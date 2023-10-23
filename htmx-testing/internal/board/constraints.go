package board

type TileConstraint struct {
	north []TileType
	east  []TileType
	south []TileType
	west  []TileType
}

var TileConstraints = map[TileType]map[Direction][]TileType{
	Sand: {
		North: []TileType{Sand, Grass, WaterSandN},
		East:  []TileType{Sand, Grass, Water},
		South: []TileType{Sand, Grass, WaterSandS},
		West:  []TileType{Sand, Grass, Water},
	},
	Grass: {
		North: []TileType{Grass, Sand, Forest},
		East:  []TileType{Grass, Sand, Forest},
		South: []TileType{Grass, Sand, Forest},
		West:  []TileType{Grass, Sand, Forest},
	},
	Forest: {
		North: []TileType{Forest, Grass},
		East:  []TileType{Forest, Grass},
		South: []TileType{Forest, Grass},
		West:  []TileType{Forest, Grass},
	},
	Water: {
		North: []TileType{Water, WaterSandS},
		East:  []TileType{Water, Sand},
		South: []TileType{Water, WaterSandN},
		West:  []TileType{Water, Sand},
	},
	WaterSandN: {
		North: []TileType{Water},
		East:  []TileType{Water, Sand, WaterSandN},
		South: []TileType{Sand},
		West:  []TileType{Water, Sand, WaterSandN},
	},
	WaterSandS: {
		North: []TileType{Sand},
		East:  []TileType{Water, Sand, WaterSandS},
		South: []TileType{Water},
		West:  []TileType{Water, Sand, WaterSandS},
	},
	WaterSandE: {},
	WaterSandW: {},

	// WaterSandTL = "Water to Sand Top Left"
	// WaterSandTR = "Water to Sand Top Right"
	// WaterSandBR = "Water to Sand Bottom Right"
	// WaterSandBL = "Water to Sand Bottom Left"
}
