package utils

import (
	"cellular-automation/model"
)

func CreateCellOnCellLocation(cellType string, placeHere model.Cords, fromGen int) model.Cell {
	return CreateCell(cellType, placeHere.GetX(), placeHere.GetY(), fromGen)
}

func CreateCell(cellType string, x int, y int, fromGen int) model.Cell {
	return model.Cell{
		CellType:       &cellType,
		X:              x,
		Y:              y,
		BornGeneration: fromGen,
	}
}

func CreateCellsCustom(xSize int, ySize int, cellType func(x int, y int) string) [][]model.Cell {
	cells := make([][]model.Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]model.Cell, xSize)

		for x := 0; x < xSize; x++ {
			ct := cellType(x, y)
			cells[y][x] = model.Cell{
				CellType:       &ct,
				X:              x,
				Y:              y,
				BornGeneration: 0,
			}
		}
	}

	return cells
}

// probably don't need to set generation number here but it doesn't hurt
func CreateCells(xSize int, ySize int, generationNum int) [][]model.Cell {
	cells := make([][]model.Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]model.Cell, xSize)

		for x := 0; x < xSize; x++ {
			cells[y][x] = model.Cell{
				CellType:       nil,
				X:              x,
				Y:              y,
				BornGeneration: generationNum,
			}
		}
	}

	return cells
}
