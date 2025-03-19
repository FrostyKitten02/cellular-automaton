package game

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"errors"
)

const ALIVE_CELL = "ALIVE"
const DEAD_CELL = "DEAD"

type NeighbourCounts struct {
	alive int
	dead  int
}

type Conway struct {
	Grid model.Grid
	Rule string
}

type ConwayRule struct {
	born    []int //number of alive neighbours a dead cell needs to be born
	survive []int //number of alive neighbours an alive cell needs to survive
}

func (c *Conway) EditGrid(grid model.Grid) {
	c.Grid = grid
}

func (c *Conway) GetGrid() *model.Grid {
	return &c.Grid
}

func (c *Conway) NextGeneration() error {
	rule := parseStringRule(c.Rule)
	nextGen := utils.CreateCells(c.Grid.XSize, c.Grid.YSize)
	for x := 0; x < c.Grid.XSize; x++ {
		for y := 0; y < c.Grid.YSize; y++ {
			counts := countNeighbours(c.Grid, x, y)
			cell := utils.GetCellFromGrid(c.Grid, x, y)
			if *cell.CellType == ALIVE_CELL {
				if ruleApplies(counts, rule.survive) {
					nextGen[y][x] = utils.CreateCell(ALIVE_CELL, x, y)
					continue
				}

				nextGen[y][x] = utils.CreateCell(DEAD_CELL, x, y)
				continue
			}

			if *cell.CellType == DEAD_CELL {
				if ruleApplies(counts, rule.born) {
					nextGen[y][x] = utils.CreateCell(ALIVE_CELL, x, y)
					continue
				}

				nextGen[y][x] = utils.CreateCell(DEAD_CELL, x, y)
				continue
			}

			return errors.New("INVALID CELL TYPE GIVEN")
		}
	}

	c.Grid.Cells = nextGen
	return nil
}

func (c *Conway) Init(xSize int, ySize int) {
	cells := utils.CreateCellsCustom(xSize, ySize, func(x int, y int) string {
		return DEAD_CELL
	})

	c.Grid.XSize = xSize
	c.Grid.YSize = ySize
	c.Grid.Cells = cells
}

func countNeighbours(grid model.Grid, cellX int, cellY int) NeighbourCounts {
	neighbours := make([]*model.Cell, 8)

	neighbours[0] = utils.GetLeftNeighbour(grid, cellX, cellY)
	neighbours[1] = utils.GetRightNeighbour(grid, cellX, cellY)

	neighbours[2] = utils.GetBottomLeftNeighbour(grid, cellX, cellY)
	neighbours[3] = utils.GetBottomNeighbour(grid, cellX, cellY)
	neighbours[4] = utils.GetBottomRightNeighbour(grid, cellX, cellY)

	neighbours[5] = utils.GetTopLeftNeighbour(grid, cellX, cellY)
	neighbours[6] = utils.GetTopNeighbour(grid, cellX, cellY)
	neighbours[7] = utils.GetTopRightNeighbour(grid, cellX, cellY)

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
