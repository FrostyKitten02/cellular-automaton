package utils

import (
	"cellular-automation/model"
)

func CreateCellOnCellLocation(cellType string, placeHere model.Cords, fromGen int) model.Cell {
	return CreateCell(cellType, placeHere.GetX(), placeHere.GetY(), fromGen, 1)
}

func CreateCellOnCellLocationWithValue(cellType string, placeHere model.Cords, fromGen int, value float64) model.Cell {
	return CreateCell(cellType, placeHere.GetX(), placeHere.GetY(), fromGen, value)
}

func CreateCell(cellType string, x int, y int, fromGen int, value float64) model.Cell {
	return model.Cell{
		CellType:       &cellType,
		X:              x,
		Y:              y,
		BornGeneration: fromGen,
		Value:          value,
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
				Value:          1,
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
				Value:          1,
			}
		}
	}

	return cells
}

func CopyCells(grid model.Grid, generationNum int) [][]model.Cell {
	xSize := grid.XSize
	ySize := grid.YSize

	cells := make([][]model.Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]model.Cell, xSize)

		for x := 0; x < xSize; x++ {
			cells[y][x] = model.Cell{
				CellType:       grid.Cells[y][x].CellType,
				X:              x,
				Y:              y,
				BornGeneration: generationNum,
				Value:          grid.Cells[y][x].Value,
			}
		}
	}

	return cells
}
