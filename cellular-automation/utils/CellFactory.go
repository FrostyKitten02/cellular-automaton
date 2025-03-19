package utils

import (
	"cellular-automation/model"
)

func CreateCell(cellType string, x int, y int) model.Cell {
	return model.Cell{
		CellType: &cellType,
		X:        x,
		Y:        y,
	}
}

func CreateEmptyCell(x int, y int) *model.Cell {
	return &model.Cell{
		X: x,
		Y: y,
	}
}

func CreateCellsCustom(xSize int, ySize int, cellType func(x int, y int) string) [][]model.Cell {
	cells := make([][]model.Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]model.Cell, xSize)

		for x := 0; x < xSize; x++ {
			ct := cellType(x, y)
			cells[y][x] = model.Cell{
				CellType: &ct,
				X:        x,
				Y:        y,
			}
		}
	}

	return cells
}

func CreateCells(xSize int, ySize int) [][]model.Cell {
	cells := make([][]model.Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]model.Cell, xSize)

		for x := 0; x < xSize; x++ {
			cells[y][x] = model.Cell{
				CellType: nil,
				X:        x,
				Y:        y,
			}
		}
	}

	return cells
}
