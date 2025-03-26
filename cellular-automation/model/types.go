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
	NextGenerationCell(currentGeneration Grid, currentCell Cell, provider ElementProvider, gameInfo GameInfo, futureGen *[][]Cell)
	GetProperties() ElementProperties
}

type GameInfo struct {
	GenerationNum     int
	CurrentGeneration Grid
}

type ElementProvider interface {
	GetBurningElements() []Element
	GetBurningElementsCellTypes() []string
	IsBurningCellType(cellType string) bool
	GetFlammableElements() []Element
	GetFlammableElementosTypes() []string
	IsFlammableCellType(cellType string) bool
}

type ElementProperties struct {
	Flammable bool //if element burns
	Burning   bool //if element burns other elements
}

type Cords interface {
	GetX() int
	GetY() int
}

type Cell struct {
	CellType       *string `json:"cellType"`
	X              int     `json:"x"`
	Y              int     `json:"y"`
	BornGeneration int     `json:"bornGeneration"`
	Value          float64 `json:"value"`
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
