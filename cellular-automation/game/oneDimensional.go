package game

import (
	"cellular-automation/model"
)

type OneDimensional struct {
	Grid model.Grid
	Rule int
}

func (o *OneDimensional) NextGeneration() error {
	parsedRule := ParseRule(o.Rule)
	//iterate iver Grid which has Grid.Cells of type [][]model.Cell store generation in next row of cells
	for gen := 1; gen < o.Grid.YSize; gen++ {
		nextGen := nextRow(o.Grid.Cells[gen-1], gen, parsedRule)
		o.Grid.Cells[gen] = nextGen
	}

	return nil
}

func (o *OneDimensional) GetGrid() *model.Grid {
	return &o.Grid
}

func (o *OneDimensional) Init(xSize int, ySize int) {
	//don't have to do anything?
}

func (o *OneDimensional) EditGrid(grid model.Grid) {
	o.Grid = grid
}

func (o *OneDimensional) GetElementProvider() model.ElementProvider {
	panic("implement me")
}

func nextRow(currentGen []model.Cell, nextGen int, rule map[[3]string]string) []model.Cell {
	return GenerateNextGen(currentGen, nextGen, rule)
}

func GenerateNextGen(cells []model.Cell, generation int, ruleMap map[[3]string]string) []model.Cell {
	n := len(cells)
	newCells := make([]model.Cell, n)

	for i := range cells {
		left := model.EmptyCell.String()
		if i > 0 {
			left = *cells[i-1].CellType
		}
		center := *cells[i].CellType
		right := model.EmptyCell.String()
		if i < n-1 {
			right = *cells[i+1].CellType
		}
		cellType := ApplyRule(left, center, right, ruleMap)
		newCells[i] = model.Cell{
			X:              cells[i].X,
			Y:              generation,
			BornGeneration: generation,
			CellType:       &cellType,
		}
	}
	return newCells
}

func ApplyRule(left string, center string, right string, ruleMap map[[3]string]string) string {
	return ruleMap[[3]string{left, center, right}]
}

func ParseRule(rule int) map[[3]string]string {
	ruleMap := make(map[[3]string]string)
	alive := model.WallCell.String()
	dead := model.EmptyCell.String()
	for i := 0; i < 8; i++ {
		for i := 0; i < 8; i++ {
			var left, center, right string

			if (i>>2)&1 == 1 {
				left = alive
			} else {
				left = dead
			}

			if (i>>1)&1 == 1 {
				center = alive
			} else {
				center = dead
			}

			if i&1 == 1 {
				right = alive
			} else {
				right = dead
			}

			pattern := [3]string{left, center, right}

			if (rule>>i)&1 == 1 {
				ruleMap[pattern] = alive
			} else {
				ruleMap[pattern] = dead
			}
		}

		return ruleMap
	}
	return ruleMap
}
