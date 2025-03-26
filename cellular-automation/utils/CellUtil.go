package utils

import (
	"cellular-automation/model"
)

func AppendCellInArr(newCell *model.Cell, oldCell *model.Cell, cells *[][]model.Cell) {
	if oldCell != nil {
		(*cells)[oldCell.GetY()][oldCell.GetX()] = *oldCell
	}

	if newCell != nil {
		(*cells)[newCell.GetY()][newCell.GetX()] = *newCell
	}
}

func AnyBurningNeighbours(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider) bool {
	x := currentCell.GetX()
	y := currentCell.GetY()

	left := GetLeftNeighbour(currentGeneration, x, y)
	if left != nil && provider.IsBurningCellType(*left.CellType) {
		return true
	}
	right := GetRightNeighbour(currentGeneration, x, y)
	if right != nil && provider.IsBurningCellType(*right.CellType) {
		return true
	}

	bottomLeft := GetBottomLeftNeighbour(currentGeneration, x, y)
	if bottomLeft != nil && provider.IsBurningCellType(*bottomLeft.CellType) {
		return true
	}
	bottom := GetBottomNeighbour(currentGeneration, x, y)
	if bottom != nil && provider.IsBurningCellType(*bottom.CellType) {
		return true
	}
	bottomRight := GetBottomRightNeighbour(currentGeneration, x, y)
	if bottomRight != nil && provider.IsBurningCellType(*bottomRight.CellType) {
		return true
	}

	topLeft := GetTopLeftNeighbour(currentGeneration, x, y)
	if topLeft != nil && provider.IsBurningCellType(*topLeft.CellType) {
		return true
	}
	top := GetTopNeighbour(currentGeneration, x, y)
	if top != nil && provider.IsBurningCellType(*top.CellType) {
		return true
	}
	topRight := GetTopRightNeighbour(currentGeneration, x, y)
	if topRight != nil && provider.IsBurningCellType(*topRight.CellType) {
		return true
	}

	return false
}

// LEFT AND RIGHT NEIGHBOURS
func GetLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY)
}

func GetRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY)
}

// BOTTOM NEIGHBOURS
func GetBottomLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY+1)
}

func GetBottomNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX, cellY+1)
}

func GetBottomRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY+1)
}

// TOP NEIGHBOURS
func GetTopLeftNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX-1, cellY-1)
}

func GetTopNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX, cellY-1)
}

func GetTopRightNeighbour(grid model.Grid, cellX int, cellY int) *model.Cell {
	return GetCellFromGrid(grid, cellX+1, cellY-1)
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
