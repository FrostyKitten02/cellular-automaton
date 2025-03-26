package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
)

var Sand = sand{
	cellType: model.SandCell,
	properties: model.ElementProperties{
		Flammable: false,
		Burning:   false,
	},
}

type sand struct {
	cellType   model.CellType
	properties model.ElementProperties
}

func (s *sand) GetProperties() model.ElementProperties {
	return s.properties
}

func (s *sand) GetCellType() model.CellType {
	return s.cellType
}

func (s *sand) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo, futureGen *[][]model.Cell) {
	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottom != nil && (*bottom.CellType == model.EmptyCell.String() || *bottom.CellType == model.Water.String()) {
		nextLocation := utils.CreateCellOnCellLocation(model.SandCell.String(), bottom, currentCell.BornGeneration)
		oldLocation := utils.CreateCellOnCellLocationWithValue(*bottom.CellType, &currentCell, gameInfo.GenerationNum, bottom.Value)
		utils.AppendCellInArr(&nextLocation, &oldLocation, futureGen)
		return
	}

	bottomLeft := utils.GetBottomLeftNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottomLeft != nil && (*bottomLeft.CellType == model.EmptyCell.String() || *bottomLeft.CellType == model.Water.String()) {
		nextLocation := utils.CreateCellOnCellLocation(model.SandCell.String(), bottomLeft, currentCell.BornGeneration)
		oldLocation := utils.CreateCellOnCellLocationWithValue(*bottomLeft.CellType, &currentCell, gameInfo.GenerationNum, bottomLeft.Value)
		utils.AppendCellInArr(&nextLocation, &oldLocation, futureGen)
		return
	}

	bottomRight := utils.GetBottomRightNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottomRight != nil && (*bottomRight.CellType == model.EmptyCell.String() || *bottomRight.CellType == model.Water.String()) {
		nextLocation := utils.CreateCellOnCellLocation(model.SandCell.String(), bottomRight, currentCell.BornGeneration)
		oldLocation := utils.CreateCellOnCellLocationWithValue(*bottomRight.CellType, &currentCell, gameInfo.GenerationNum, bottomRight.Value)
		utils.AppendCellInArr(&nextLocation, &oldLocation, futureGen)
		return
	}

	sameLocation := utils.CreateCellOnCellLocation(model.SandCell.String(), &currentCell, currentCell.BornGeneration)
	utils.AppendCellInArr(&sameLocation, nil, futureGen)
}
