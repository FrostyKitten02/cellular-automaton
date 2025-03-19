package utils

import (
	"cellular-automation/model"
)

// LEFT AND RIGHT NEIGHBOURS
func GetLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY)
}

func GetRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY)
}

// BOTTOM NEIGHBOURS
func GetBottomLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY-1)
}

func GetBottomNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX, cellY-1)
}

func GetBottomRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY-1)
}

// TOP NEIGHBOURS
func GetTopLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY+1)
}

func GetTopNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX, cellY+1)
}

func GetTopRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY+1)
}

func GetCellFromGrid(grid model.Grid, x int, y int) *model.Cell {
	maxY := grid.YSize - 1
	maxX := grid.XSize - 1

	if y > maxY || y < 0 {
		return nil
	}

	if x > maxX || x < 0 {
		return nil
	}

	return &grid.Cells[y][x]
}
