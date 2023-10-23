package board

type TileConstraint struct {
	tileType TileType
	north    []TileType
	east     []TileType
	south    []TileType
	west     []TileType
}

var TileConstraints = []TileConstraint{
	{
		tileType: Sand,
		north:    []TileType{Sand, Grass, Water},
		east:     []TileType{Sand, Grass, Water},
		south:    []TileType{Sand, Grass, Water},
		west:     []TileType{Sand, Grass, Water},
	},
	{
		tileType: Grass,
		north:    []TileType{Grass, Sand, Forest},
		east:     []TileType{Grass, Sand, Forest},
		south:    []TileType{Grass, Sand, Forest},
		west:     []TileType{Grass, Sand, Forest},
	},
	{
		tileType: Forest,
		north:    []TileType{Forest, Grass},
		east:     []TileType{Forest, Grass},
		south:    []TileType{Forest, Grass},
		west:     []TileType{Forest, Grass},
	},
	{
		tileType: Water,
		north:    []TileType{Water, Sand},
		east:     []TileType{Water, Sand},
		south:    []TileType{Water, Sand},
		west:     []TileType{Water, Sand},
	},
}
