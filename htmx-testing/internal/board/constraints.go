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
		East:  []TileType{Sand, Grass, WaterSandE},
		South: []TileType{Sand, Grass, WaterSandS},
		West:  []TileType{Sand, Grass, WaterSandW},
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
		East:  []TileType{Water, WaterSandW},
		South: []TileType{Water, WaterSandN},
		West:  []TileType{Water, WaterSandE},
	},
	WaterSandN: {
		North: []TileType{Water},
		East:  []TileType{WaterSandN, WaterSandTL},
		South: []TileType{Sand},
		West:  []TileType{WaterSandN, WaterSandTR},
	},
	WaterSandS: {
		North: []TileType{Sand},
		East:  []TileType{WaterSandS, WaterSandBL},
		South: []TileType{Water},
		West:  []TileType{WaterSandS, WaterSandBR},
	},
	WaterSandE: {
		North: []TileType{WaterSandE, WaterSandBR},
		East:  []TileType{Water},
		South: []TileType{WaterSandE, WaterSandTR},
		West:  []TileType{Sand},
	},
	WaterSandW: {
		North: []TileType{WaterSandW, WaterSandBL},
		East:  []TileType{Sand},
		South: []TileType{WaterSandW, WaterSandTL},
		West:  []TileType{Water},
	},
	WaterSandTL: {
		North: []TileType{WaterSandW},
		East:  []TileType{Sand},
		South: []TileType{Sand},
		West:  []TileType{WaterSandN},
	},
	WaterSandTR: {
		North: []TileType{WaterSandE},
		East:  []TileType{WaterSandN},
		South: []TileType{Sand},
		West:  []TileType{Sand},
	},
	WaterSandBR: {
		North: []TileType{Sand},
		East:  []TileType{WaterSandS},
		South: []TileType{WaterSandE},
		West:  []TileType{Sand},
	},
	WaterSandBL: {
		North: []TileType{Sand},
		East:  []TileType{Sand},
		South: []TileType{WaterSandW},
		West:  []TileType{WaterSandS},
	},
}
