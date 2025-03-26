package model

type CellType string

const (
	//CONWAY CELLS START
	WallCell  CellType = "WALL"  //this is alive cell in normal conway game of life
	EmptyCell CellType = "EMPTY" //this is dead cell in normal conway game of life
	//CONWAY CELLS END
	//CUSTOM CELLS
	SandCell   CellType = "SAND"
	WoodCell   CellType = "WOOD"
	FireCell   CellType = "FIRE"
	DarkSmoke  CellType = "DARK_SMOKE"
	WhiteSmoke CellType = "WHITE_SMOKE"
	Water      CellType = "WATER"
)

func (c CellType) String() string {
	return string(c)
}
