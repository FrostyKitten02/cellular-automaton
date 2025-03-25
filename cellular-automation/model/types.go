package model

type Game interface {
	NextGeneration() error
	GetGrid() *Grid
	Init(xSize int, ySize int)
	EditGrid(Grid)
	GetElementProvider() ElementProvider
}

type Element interface {
	GetCellType() CellType
	NextGenerationCell(currentGeneration Grid, currentCell Cell, provider ElementProvider) Cell
	GetProperties() ElementProperties
}

type ElementProvider interface {
	GetBurningElements() []Element
	GetBurningElementsCellTypes() []string
	IsBurningCellType(cellType string) bool
}

type ElementProperties struct {
	Flameable bool //if element burns
	Burning   bool //if element burns other elements
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
