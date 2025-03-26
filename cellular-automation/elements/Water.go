package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"log"
)

var Water = water{
	cellType:    model.Water,
	maxCapacity: 3,
	flowRate:    0.5,
	minPressure: 1.5,
	properties: model.ElementProperties{
		Flammable: false,
		Burning:   false,
	},
}

type water struct {
	cellType    model.CellType
	properties  model.ElementProperties
	maxCapacity float64 //soft limit
	minPressure float64 //when cell is allowed to move up
	flowRate    float64
}

func (w *water) GetProperties() model.ElementProperties {
	return w.properties
}

func (w *water) GetCellType() model.CellType {
	return w.cellType
}

func (w *water) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo, futureGen *[][]model.Cell) {
	top := utils.GetTopNeighbour(currentGeneration, currentCell.X, currentCell.Y)

	//don't move if sand above, bcs it will get replaced
	if top != nil && *top.CellType == model.SandCell.String() {
		oldCellValue := getValue(currentCell, *futureGen, currentGeneration)

		oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
		if oldCell.Value <= 0 {
			//changing type to empty if value is 0, when cell gets empty
			oldCell = utils.CreateCellOnCellLocationWithValue(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
		}
		return
	}

	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())

	//move down
	if bottom != nil && (*bottom.CellType == model.EmptyCell.String() || *bottom.CellType == model.Water.String()) {
		log.Print("Water attempting to move down!")
		newCellValue := getValue(*bottom, *futureGen, currentGeneration)
		oldCellValue := getValue(currentCell, *futureGen, currentGeneration)

		log.Print("Old cell before value:", oldCellValue)
		log.Print("New cell after value:", newCellValue)

		canMoveDown := newCellValue < w.maxCapacity

		if !canMoveDown {
			oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			if oldCell.Value <= 0 {
				//changing type to empty if value is 0, when cell gets empty
				oldCell = utils.CreateCellOnCellLocationWithValue(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			}

			log.Print("Water moved down!!")
			utils.AppendCellInArr(nil, &oldCell, futureGen)
			return
		}

		if oldCellValue < w.flowRate {
			newCellValue += oldCellValue
			oldCellValue -= oldCellValue
		} else {
			newCellValue += w.flowRate
			oldCellValue -= w.flowRate
		}

		if canMoveDown {
			log.Print("Old cell value:", oldCellValue)
			log.Print("New cell value:", newCellValue)

			oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			if oldCell.Value <= 0 {
				//changing type to empty if value is 0, when cell gets empty
				oldCell = utils.CreateCellOnCellLocationWithValue(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			}

			bottomCell := utils.CreateCellOnCellLocationWithValue(w.cellType.String(), bottom, gameInfo.GenerationNum, newCellValue)
			log.Print("Water moved down!!")
			utils.AppendCellInArr(&bottomCell, &oldCell, futureGen)
			return
		}
	}

	//should go left and right
	if bottom != nil && (*bottom.CellType != model.EmptyCell.String() || *bottom.CellType == model.Water.String()) {
		left := utils.GetLeftNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
		oldCellValue := getValue(currentCell, *futureGen, currentGeneration)
		leftCellValue := getValue(*left, *futureGen, currentGeneration)

		futureGrid := model.Grid{
			Cells: *futureGen,
			XSize: currentGeneration.XSize,
			YSize: currentGeneration.YSize,
		}
		futLeft := utils.GetCellFromGrid(futureGrid, left.GetX(), left.GetY())
		wentLeft := false
		if (*left.CellType == model.EmptyCell.String() || *left.CellType == model.Water.String()) && (futLeft.CellType == nil || *futLeft.CellType != model.SandCell.String()) {
			if oldCellValue < w.flowRate {
				if leftCellValue >= w.maxCapacity && leftCellValue > oldCellValue {
					//Then should check to move up!!
					wentLeft = false
				} else {
					if leftCellValue > oldCellValue {
						wentLeft = false
					} else {
						wentLeft = true
						leftCellValue += oldCellValue
						oldCellValue -= oldCellValue
					}
				}
			} else {
				if leftCellValue >= w.maxCapacity && leftCellValue > oldCellValue {
					//Then should check to move up!!
					wentLeft = false
				} else {
					if leftCellValue > oldCellValue {
						wentLeft = false
					} else {
						wentLeft = true
						leftCellValue += w.flowRate
						oldCellValue -= w.flowRate
					}
				}
			}

			if wentLeft {
				newLeftCell := utils.CreateCellOnCellLocationWithValue(w.cellType.String(), left, gameInfo.GenerationNum, leftCellValue)
				log.Print("Water moved left!!")
				utils.AppendCellInArr(&newLeftCell, nil, futureGen)
			}
		}

		right := utils.GetRightNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
		rightCellValue := getValue(*right, *futureGen, currentGeneration)

		futRight := utils.GetCellFromGrid(futureGrid, right.GetX(), right.GetY())
		wentRight := false
		if (*right.CellType == model.EmptyCell.String() || *right.CellType == model.Water.String()) && (futRight.CellType == nil || *futRight.CellType != model.SandCell.String()) {
			if oldCellValue < w.flowRate {
				if rightCellValue >= w.maxCapacity && rightCellValue > oldCellValue {
					wentRight = false
				} else if rightCellValue <= oldCellValue {
					wentRight = true
					rightCellValue += oldCellValue
					oldCellValue -= oldCellValue
				}
			} else {
				if oldCellValue >= w.maxCapacity && rightCellValue > oldCellValue {
					wentRight = false
				} else if rightCellValue <= oldCellValue {
					wentRight = true
					rightCellValue += w.flowRate
					oldCellValue -= w.flowRate
				}
			}

			if wentRight {
				newRightCell := utils.CreateCellOnCellLocationWithValue(w.cellType.String(), right, gameInfo.GenerationNum, rightCellValue)
				log.Print("Water moved right!!")
				utils.AppendCellInArr(&newRightCell, nil, futureGen)
			}
		}

		if wentRight || wentLeft {
			log.Println("Updating prev water", oldCellValue)
			oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			if oldCell.Value <= 0 {
				//changing type to empty if value is 0, when cell gets empty
				oldCell = utils.CreateCellOnCellLocationWithValue(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
			}
			utils.AppendCellInArr(&oldCell, &oldCell, futureGen)
		}
	}

	//pressure simulation
	if top != nil && (*top.CellType == model.EmptyCell.String() || *top.CellType == model.Water.String()) {
		//TODO move water up!!
		oldCellValue := getValue(currentCell, *futureGen, currentGeneration)
		if oldCellValue <= w.minPressure {
			//No pressure simulation if pressure too small
			oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, currentCell.BornGeneration, oldCellValue)
			utils.AppendCellInArr(nil, &oldCell, futureGen)
			return
		}
		topCellValue := getValue(*top, *futureGen, currentGeneration)

		log.Print("Water before move up!!", oldCellValue, topCellValue)

		//don't move if top water has more water!!
		if oldCellValue < topCellValue {
			oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, currentCell.BornGeneration, oldCellValue)
			utils.AppendCellInArr(nil, &oldCell, futureGen)
			return
		}

		if oldCellValue < w.flowRate {
			if topCellValue >= w.maxCapacity {
				oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, currentCell.BornGeneration, oldCellValue)
				utils.AppendCellInArr(nil, &oldCell, futureGen)
				return
			}
			topCellValue += oldCellValue
			oldCellValue -= oldCellValue
		} else {
			if topCellValue >= w.maxCapacity {
				oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, currentCell.BornGeneration, oldCellValue)
				utils.AppendCellInArr(nil, &oldCell, futureGen)
				return
			}
			topCellValue += w.flowRate
			oldCellValue -= w.flowRate
		}

		topCell := utils.CreateCellOnCellLocationWithValue(w.cellType.String(), top, gameInfo.GenerationNum, topCellValue)
		log.Print("Water moved up!!", oldCellValue, topCellValue)
		oldCell := utils.CreateCellOnCellLocationWithValue(w.GetCellType().String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
		if oldCell.Value <= 0 {
			//changing type to empty if value is 0, when cell gets empty
			oldCell = utils.CreateCellOnCellLocationWithValue(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum, oldCellValue)
		}

		utils.AppendCellInArr(&topCell, &oldCell, futureGen)
		return
	}

	//if we cannot move anywhere stay on same spot
	oldCell := utils.CreateCellOnCellLocationWithValue(w.cellType.String(), &currentCell, gameInfo.GenerationNum, getValue(currentCell, *futureGen, currentGeneration))
	utils.AppendCellInArr(nil, &oldCell, futureGen)
}

func getValue(cell model.Cell, future [][]model.Cell, curr model.Grid) float64 {
	currCell := utils.GetCellFromGrid(curr, cell.GetX(), cell.GetY())

	futureGrid := model.Grid{
		Cells: future,
		XSize: curr.XSize,
		YSize: curr.YSize,
	}
	log.Println("Created temp future grid")
	futureCell := utils.GetCellFromGrid(futureGrid, cell.GetX(), cell.GetY())
	log.Println("Got future cell:", futureCell)
	if futureCell != nil && futureCell.CellType != nil && *futureCell.CellType == model.Water.String() {
		log.Println("Found water cell:", futureCell)
		return futureCell.Value
	}

	if currCell != nil && *currCell.CellType == model.Water.String() {
		log.Println("Found water cell 2:", currCell)
		return currCell.Value
	}

	log.Println("Did not find water cell returning 0")
	return 0
}
