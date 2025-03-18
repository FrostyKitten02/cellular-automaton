package main

import (
	"errors"
)

const ALIVE_CELL = "ALIVE"

const DEAD_CELL = "DEAD"

type NeighbourCounts struct {
	alive int
	dead  int
}

func initialConway(xSize int, ySize int) [][]Cell {
	return createCellsCustom(xSize, ySize, func(x int, y int) string {
		return DEAD_CELL
	})
}

func nexGeneration(grid Grid) ([][]Cell, error) {
	nextGen := createCells(grid.XSize, grid.YSize)
	//create when exactly 3 neighbours alive
	//keep alive 2 or 3 neighbours
	for x := 0; x < grid.XSize; x++ {
		for y := 0; y < grid.YSize; y++ {
			counts := countNeighbours(grid, x, y)
			cell := getCellFromGrid(grid, x, y)
			if *cell.CellType == ALIVE_CELL {
				if counts.alive == 2 || counts.alive == 3 {
					nextGen[y][x] = createCell(ALIVE_CELL, x, y)
					continue
				}

				nextGen[y][x] = createCell(DEAD_CELL, x, y)
				continue
			}

			if *cell.CellType == DEAD_CELL {
				if counts.alive == 3 {
					nextGen[y][x] = createCell(ALIVE_CELL, x, y)
					continue
				}

				nextGen[y][x] = createCell(DEAD_CELL, x, y)
				continue
			}

			return nil, errors.New("INVALID CELL TYPE GIVEN")
		}
	}

	return nextGen, nil
}

func countNeighbours(grid Grid, cellX int, cellY int) NeighbourCounts {
	neighbours := make([]*Cell, 8)

	neighbours[0] = getLeftNeighbour(grid, cellX, cellY)
	neighbours[1] = getRightNeighbour(grid, cellX, cellY)

	neighbours[2] = getBottomLeftNeighbour(grid, cellX, cellY)
	neighbours[3] = getBottomNeighbour(grid, cellX, cellY)
	neighbours[4] = getBottomRightNeighbour(grid, cellX, cellY)

	neighbours[5] = getTopLeftNeighbour(grid, cellX, cellY)
	neighbours[6] = getTopNeighbour(grid, cellX, cellY)
	neighbours[7] = getTopRightNeighbour(grid, cellX, cellY)

	res := NeighbourCounts{
		alive: 0,
		dead:  0,
	}
	for _, cell := range neighbours {
		if cell == nil {
			continue
		}
		if *cell.CellType == ALIVE_CELL {
			res.alive++
		}

		if *cell.CellType == DEAD_CELL {
			res.dead++
		}
	}

	return res
}
