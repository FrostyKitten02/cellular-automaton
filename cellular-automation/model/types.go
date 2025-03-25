package model

type Game interface {
	NextGeneration() error
	GetGrid() *Grid
	Init(xSize int, ySize int)
	EditGrid(Grid)
}

type Element interface {
	GetCellType() CellType
	NextGenerationCell(currentGeneration Grid, currentCell Cell) Cell
}

type Cords interface {
	GetX() int
	GetY() int
}

type Cell struct {
	CellType *string `json:"cellType"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
}

func (c *Cell) GetX() int {
	return c.X
}

func (c *Cell) GetY() int {
	return c.Y
}

type Grid struct {
	Cells      [][]Cell `json:"Cells"`
	XSize      int      `json:"xSize"`
	YSize      int      `json:"ySize"`
	InProgress bool     `json:"inProgress"`
}
