package main

// LEFT AND RIGHT NEIGHBOURS
func getLeftNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX-1, cellY)
}

func getRightNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX+1, cellY)
}

// BOTTOM NEIGHBOURS
func getBottomLeftNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX-1, cellY-1)
}

func getBottomNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX, cellY-1)
}

func getBottomRightNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX+1, cellY-1)
}

// TOP NEIGHBOURS
func getTopLeftNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX-1, cellY+1)
}

func getTopNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX, cellY+1)
}

func getTopRightNeighbour(grid Grid, cellX int, cellY int) *Cell {
	return getCellFromGrid(grid, cellX+1, cellY+1)
}

func getCellFromGrid(grid Grid, x int, y int) *Cell {
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
