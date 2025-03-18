package main

func createCell(cellType string, x int, y int) Cell {
	return Cell{
		CellType: &cellType,
		X:        x,
		Y:        y,
	}
}

func createEmptyCell(x int, y int) *Cell {
	return &Cell{
		X: x,
		Y: y,
	}
}

func createCellsCustom(xSize int, ySize int, cellType func(x int, y int) string) [][]Cell {
	cells := make([][]Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]Cell, xSize)

		for x := 0; x < xSize; x++ {
			ct := cellType(x, y)
			cells[y][x] = Cell{
				CellType: &ct,
				X:        x,
				Y:        y,
			}
		}
	}

	return cells
}

func createCells(xSize int, ySize int) [][]Cell {
	cells := make([][]Cell, ySize)
	for y := 0; y < ySize; y++ {
		cells[y] = make([]Cell, xSize)

		for x := 0; x < xSize; x++ {
			cells[y][x] = Cell{
				CellType: nil,
				X:        x,
				Y:        y,
			}
		}
	}

	return cells
}
