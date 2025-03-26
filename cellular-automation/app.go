package main

import (
	"cellular-automation/game"
	"cellular-automation/model"
	"cellular-automation/utils"
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"strconv"
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
	log.Print("Step called")
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
	log.Print("Simulate called")
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
	log.Print("Stop simulation called")
	if a.simulationCancelFunc != nil {
		a.simulationCancelFunc()
		a.simulationCancelFunc = nil
	}

	return *a.game.GetGrid()
}

func (a *App) ResetGrid() model.Grid {
	log.Print("Reset grid called")
	return model.Grid{}
}

func (a *App) Init(xSize int, ySize int, gameMode string, options map[string]string) (*model.Grid, error) {
	log.Print("Init called for ", gameMode)
	if a.simulationCancelFunc != nil {
		return nil, errors.New("Simulation in progress")
	}

	if gameMode == "CONWAY" {
		conwayCondition, alivePercent, err := parseConwayOpts(options)
		if err != nil {
			return nil, err
		}
		a.game = game.NewConway(conwayCondition, alivePercent)
	}

	if gameMode == "SANDBOX" {
		conwayCondition, alivePercent, err := parseConwayOpts(options)
		if err != nil {
			return nil, err
		}
		pregenCave := options["pregenCave"] == "true"
		a.game, err = game.NewSandbox(conwayCondition, alivePercent, pregenCave, xSize, ySize)
		if err != nil {
			return nil, err
		}
	}

	if gameMode == "1D" {
		ruleStr := options["rule"]
		rule, err := strconv.Atoi(ruleStr)
		if err != nil {
			return nil, err
		}

		oneDimGame := game.OneDimensional{
			Grid: model.Grid{
				Cells: utils.CreateOneDimensionalGrid(xSize, ySize),
				XSize: xSize,
				YSize: ySize,
			},
			Rule: rule,
		}

		a.game = &oneDimGame
		return a.game.GetGrid(), nil
	}

	a.game.Init(xSize, ySize)
	return a.game.GetGrid(), nil
}

func (a *App) EditGrid(grid model.Grid) model.Grid {
	//TODO should validate grid size
	log.Print("Edit called")
	a.game.EditGrid(grid)
	return *a.game.GetGrid()
}

func parseConwayOpts(options map[string]string) (string, int, error) {
	aliveOption := options["alivePercent"]
	alivePercent := 0
	if aliveOption != "" {
		parseRes, err := strconv.Atoi(aliveOption)
		if err != nil {
			return "", 0, errors.New("Invalid value for 'alivePercent': " + options["alivePercent"])
		}
		alivePercent = parseRes
	}

	conwayCondition := options["conwayCondition"]
	if conwayCondition == "" {
		conwayCondition = "B3/S23" //default value for conway if no condition is sent
	}

	return conwayCondition, alivePercent, nil
}
