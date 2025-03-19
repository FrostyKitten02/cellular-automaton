package model

type Game interface {
	NextGeneration() error
	GetGrid() *Grid
	Init(xSize int, ySize int)
	EditGrid(Grid)
}

type Cell struct {
	CellType *string `json:"cellType"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
}

type Grid struct {
	Cells      [][]Cell `json:"Cells"`
	XSize      int      `json:"xSize"`
	YSize      int      `json:"ySize"`
	InProgress bool     `json:"inProgress"`
}
