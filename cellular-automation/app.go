package main

import (
	"context"
)

// App struct
type App struct {
	ctx  context.Context
	grid Grid
}

type Cell struct {
	CellType *string `json:"cellType"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
}

type Grid struct {
	Cells [][]Cell `json:"Cells"`
	XSize int      `json:"xSize"`
	YSize int      `json:"ySize"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Step() *Grid {
	cells, err := nexGeneration(a.grid)
	if err != nil {
		return nil
	}
	a.grid.Cells = cells

	return &a.grid
}

func (a *App) Simulate() Grid {
	return Grid{}
}

func (a *App) ResetGrid() Grid {
	return Grid{}
}

func (a *App) InitGrid(xSize int, ySize int) Grid {
	//grid := Grid{
	//	Cells: createCells(xSize, ySize),
	//	XSize: xSize,
	//	YSize: ySize,
	//}

	grid := Grid{
		Cells: initialConway(xSize, ySize),
		XSize: xSize,
		YSize: ySize,
	}

	a.grid = grid
	return grid
}
