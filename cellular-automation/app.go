package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx                  context.Context
	simulationCancelFunc context.CancelFunc
	grid                 Grid
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
	a.stepInternal()
	return &a.grid
}

func (a *App) stepInternal() {
	cells, err := nexGeneration(a.grid)
	if err != nil {
		a.grid.Cells = nil
	}
	a.grid.Cells = cells
}

func (a *App) Simulate() {
	var streamCtx context.Context
	streamCtx, a.simulationCancelFunc = context.WithCancel(a.ctx) // Create a cancelable context

	go func() {
		for {
			select {
			case <-streamCtx.Done():
				runtime.LogInfo(a.ctx, "Stream stopped by client")
				return
			default:
				a.stepInternal()
				runtime.EventsEmit(a.ctx, "simulation_stream", a.grid)
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
}

func (a *App) StopSimulation() Grid {
	if a.simulationCancelFunc != nil {
		a.simulationCancelFunc()
		a.simulationCancelFunc = nil
	}

	return a.grid
}

func (a *App) ResetGrid() Grid {
	return Grid{}
}

func (a *App) InitGrid(xSize int, ySize int) Grid {
	grid := Grid{
		Cells:      initialConway(xSize, ySize),
		XSize:      xSize,
		YSize:      ySize,
		InProgress: false,
	}

	a.grid = grid
	return grid
}

func (a *App) EditGrid(grid Grid) Grid {
	//TODO should validate grid size
	a.grid = grid
	return a.grid
}
