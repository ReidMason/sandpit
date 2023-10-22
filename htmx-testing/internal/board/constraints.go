package board

type TileConstraint struct {
	tileType TileType
	north    TileType
	east     TileType
	south    TileType
	west     TileType
}

var TileConstraints = []TileConstraint{
	// Sand
	{
		tileType: Sand,
		north:    Sand,
		east:     Sand,
		south:    Sand,
		west:     Sand,
	},
	{
		tileType: Sand,
		north:    Water,
		east:     Water,
		south:    Water,
		west:     Water,
	},
	{
		tileType: Sand,
		north:    Grass,
		east:     Grass,
		south:    Grass,
		west:     Grass,
	},
	// Grass
	{
		tileType: Grass,
		north:    Grass,
		east:     Grass,
		south:    Grass,
		west:     Grass,
	},
	// Forest
	{
		tileType: Forest,
		north:    Forest,
		east:     Forest,
		south:    Forest,
		west:     Forest,
	},
	{
		tileType: Forest,
		north:    Grass,
		east:     Grass,
		south:    Grass,
		west:     Grass,
	},
	// Water
	{
		tileType: Water,
		north:    Water,
		east:     Water,
		south:    Water,
		west:     Water,
	},
	// // Water to sand
	// {
	// 	tileType: WaterSandTL,
	// 	north:    Water,
	// 	east:     Sand,
	// 	south:    Sand,
	// 	west:     Water,
	// },
	// {
	// 	tileType: WaterSandTR,
	// 	north:    Water,
	// 	east:     Water,
	// 	south:    Sand,
	// 	west:     Sand,
	// },
	// {
	// 	tileType: WaterSandBR,
	// 	north:    Sand,
	// 	east:     Water,
	// 	south:    Water,
	// 	west:     Sand,
	// },
	// {
	// 	tileType: WaterSandBL,
	// 	north:    Sand,
	// 	east:     Sand,
	// 	south:    Water,
	// 	west:     Water,
	// },
	// {
	// 	tileType: WaterSandN,
	// 	north:    Water,
	// 	east:     WaterSandTR,
	// 	south:    Sand,
	// 	west:     WaterSandTL,
	// },
	// {
	// 	tileType: WaterSandS,
	// 	north:    Sand,
	// 	east:     WaterSandBR,
	// 	south:    Water,
	// 	west:     WaterSandBL,
	// },
	// {
	// 	tileType: WaterSandE,
	// 	north:    WaterSandTR,
	// 	east:     Water,
	// 	south:    WaterSandBR,
	// 	west:     Sand,
	// },
	// {
	// 	tileType: WaterSandW,
	// 	north:    WaterSandTL,
	// 	east:     Sand,
	// 	south:    WaterSandBL,
	// 	west:     Water,
	// },
}
