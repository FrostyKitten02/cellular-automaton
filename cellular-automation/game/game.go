package game

import (
	"cellular-automation/elements"
	"cellular-automation/model"
	"cellular-automation/utils"
	"errors"
	"log"
	"math/rand"
)

type NeighbourCounts struct {
	alive int
	dead  int
}

type GenerationArgs struct {
	alivePercent    int
	elementsPercent map[string]int //map of cellType and their chance of spawning in empty spaces
}

type SandboxGame struct {
	Grid             model.Grid
	Rule             ConwayRule
	elements         *[]model.Element
	genArgs          GenerationArgs
	elementsProvider model.ElementProvider
	generationNum    int
	cavePregen       bool
}

type ConwayRule struct {
	born    []int //number of alive neighbours a dead cell needs to be born
	survive []int //number of alive neighbours an alive cell needs to survive
}

func (c *SandboxGame) GetGameInfo() model.GameInfo {
	return model.GameInfo{
		GenerationNum:     c.generationNum,
		CurrentGeneration: c.Grid,
	}
}

func (c *SandboxGame) GetElementProvider() model.ElementProvider {
	return c.elementsProvider
}

func (c *SandboxGame) EditGrid(grid model.Grid) {
	cells := &grid.Cells

	a := 0
	b := 0
	for y := 0; y < grid.YSize; y++ {
		for x := 0; x < grid.XSize; x++ {
			if (*cells)[y][x].BornGeneration == -1 {
				log.Print("NEW CELL DETECTED", (*cells)[y][x])
				(*cells)[y][x].BornGeneration = c.generationNum
				a = y
				b = x
			}
		}
	}

	c.Grid = model.Grid{
		Cells: *cells,
		XSize: grid.XSize,
		YSize: grid.YSize,
	}
	log.Print(c.Grid.Cells[a][b])
}

func (c *SandboxGame) GetGrid() *model.Grid {
	return &c.Grid
}

func (c *SandboxGame) NextGeneration() error {
	c.generationNum = c.generationNum + 1

	gameInfo := c.GetGameInfo()
	currentGen := c.Grid
	nextGen := utils.CreateCells(currentGen.XSize, currentGen.YSize, gameInfo.GenerationNum)
	//nextGen := utils.CopyCells(c.Grid, gameInfo.GenerationNum)
	for x := 0; x < currentGen.XSize; x++ {
		for y := 0; y < currentGen.YSize; y++ {
			counts := countNeighbours(currentGen, x, y)
			cell := utils.GetCellFromGrid(currentGen, x, y)
			if *cell.CellType == model.WallCell.String() {
				//maybe we will need to check here also if any cell was already created before creating it???
				if ruleApplies(counts, c.Rule.survive) {
					nextGen[y][x] = utils.CreateCell(model.WallCell.String(), x, y, cell.BornGeneration, 1)
					continue
				}

				nextGen[y][x] = utils.CreateCell(model.EmptyCell.String(), x, y, gameInfo.GenerationNum, 1)
				continue
			}

			if *cell.CellType == model.EmptyCell.String() {
				//not replacing any blocks if block was already created!
				if nextGen[y][x].CellType != nil {
					continue
				}

				if ruleApplies(counts, c.Rule.born) {
					nextGen[y][x] = utils.CreateCell(model.WallCell.String(), x, y, gameInfo.GenerationNum, 1)
					continue
				}

				nextGen[y][x] = utils.CreateCell(model.EmptyCell.String(), x, y, cell.BornGeneration, 1)
				continue
			}

			if c.elements == nil {
				//if we get to here it means we don't have anymore cellTypes to look for so we can throw
				return errors.New("INVALID CELL TYPE GIVEN")
			}

			element := utils.FindElementForCellType(c.elements, *cell.CellType)

			//throwing exception if element is nil, this means we don't have that element implemented
			if element == nil {
				return errors.New("INVALID CELL TYPE GIVEN")
			}

			(*element).NextGenerationCell(currentGen, *cell, c.elementsProvider, gameInfo, &nextGen)
		}
	}

	for x := 0; x < currentGen.XSize; x++ {
		for y := 0; y < currentGen.YSize; y++ {
			if nextGen[y][x].CellType == nil || *nextGen[y][x].CellType == "" {
				emptyCell := model.EmptyCell.String()
				nextGen[y][x].CellType = &emptyCell
			}
		}
	}

	c.Grid.Cells = nextGen
	return nil
}

func (c *SandboxGame) Init(xSize int, ySize int) {
	cells := utils.CreateCellsCustom(xSize, ySize, func(x int, y int) string {
		//only generate walls if cave was not pre-generated
		if c.cavePregen && *utils.GetCellFromGrid(c.Grid, x, y).CellType == model.WallCell.String() {
			return model.WallCell.String()
		}

		if !c.cavePregen && rand.Intn(101) <= c.genArgs.alivePercent {
			return model.WallCell.String()
		}

		if c.elements == nil {
			return model.EmptyCell.String()
		}

		for _, element := range *c.elements {
			val := c.genArgs.elementsPercent[element.GetCellType().String()]
			if &val == nil {
				continue
			}

			if rand.Intn(101) <= val {
				return element.GetCellType().String()
			}
		}

		return model.EmptyCell.String()
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
		if cell == nil || cell.CellType == nil {
			continue
		}

		if *cell.CellType == model.WallCell.String() {
			res.alive++
		}

		if *cell.CellType == model.EmptyCell.String() {
			res.dead++
		}
	}

	return res
}

func NewConway(rule string, alivePercent int) *SandboxGame {
	parsedRule := parseStringRule(rule)
	return &SandboxGame{
		Rule: parsedRule,
		genArgs: GenerationArgs{
			alivePercent: alivePercent,
		},
	}
}

func NewSandbox(rule string, alivePercent int, pregenCave bool, xSize int, ySize int) (*SandboxGame, error) {
	parsedRule := parseStringRule(rule)

	var pregen *SandboxGame = nil
	if pregenCave {
		pregen = NewConway(rule, alivePercent)
		pregen.Init(xSize, ySize)
		//100 generations should be enough
		for i := 0; i < 100; i++ {
			err := pregen.NextGeneration()
			if err != nil {
				return nil, errors.New("Error in pregen")
			}
		}
	}

	gameElements := &[]model.Element{&elements.Sand, &elements.Wood, &elements.Fire, &elements.DarkSmoke, &elements.WhiteSmoke, &elements.Water}
	//gameElements := &[]model.Element{&elements.Sand, &elements.Wood, &elements.Fire, &elements.DarkSmoke, &elements.WhiteSmoke}
	game := &SandboxGame{
		Rule: parsedRule,
		genArgs: GenerationArgs{
			alivePercent:    alivePercent,
			elementsPercent: map[string]int{model.SandCell.String(): 0, model.WoodCell.String(): 0, model.FireCell.String(): 20},
		},
		elements:         gameElements,
		elementsProvider: model.NewElementProvider(*gameElements),
		cavePregen:       pregenCave,
	}

	if pregen != nil {
		log.Println("Pregen:", pregen)
		game.Grid = pregen.Grid
	}

	return game, nil
}
