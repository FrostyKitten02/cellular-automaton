package main

import (
	"cellular-automation/game"
	"cellular-automation/model"
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx                  context.Context
	simulationCancelFunc context.CancelFunc
	game                 model.Game
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

func (a *App) Step() (*model.Grid, error) {
	if a.simulationCancelFunc != nil {
		return nil, errors.New("Simulation in progress")
	}

	err := a.stepInternal()
	if err != nil {
		return nil, err
	}

	return a.game.GetGrid(), nil
}

func (a *App) stepInternal() error {
	err := a.game.NextGeneration()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Simulate() error {
	if a.simulationCancelFunc != nil {
		return errors.New("Simulation already in progress")
	}

	var streamCtx context.Context
	streamCtx, a.simulationCancelFunc = context.WithCancel(a.ctx) // Create a cancelable context

	go func() {
		for {
			select {
			case <-streamCtx.Done():
				runtime.LogInfo(a.ctx, "Stream stopped by client")
				return
			default:
				err := a.stepInternal()
				if err != nil {
					//TODO give client information about cancel!
					a.simulationCancelFunc()
					return
				}
				runtime.EventsEmit(a.ctx, "simulation_stream", a.game.GetGrid())
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()

	return nil
}

func (a *App) StopSimulation() model.Grid {
	if a.simulationCancelFunc != nil {
		a.simulationCancelFunc()
		a.simulationCancelFunc = nil
	}

	return *a.game.GetGrid()
}

func (a *App) ResetGrid() model.Grid {
	return model.Grid{}
}

func (a *App) Init(xSize int, ySize int) model.Grid {
	a.game = &game.Conway{
		Rule: "B678/S2345678",
	}
	a.game.Init(xSize, ySize)
	return *a.game.GetGrid()
}

func (a *App) EditGrid(grid model.Grid) model.Grid {
	//TODO should validate grid size
	a.game.EditGrid(grid)
	return *a.game.GetGrid()
}
