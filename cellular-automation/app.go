package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

type Cell struct {
}

type Grid struct {
	Cells [][]Cell `json:"Cells"`
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

func (a *App) Step() Grid {
	return Grid{}
}

func (a *App) Simulate() Grid {
	return Grid{}
}

func (a *App) ResetGrid() Grid {
	return Grid{}
}

func (a *App) TestFunc(c Cell) {

}
