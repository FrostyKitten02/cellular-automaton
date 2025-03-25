package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
)

var Sand = sand{
	cellType: model.SandCell,
	properties: model.ElementProperties{
		Flameable: false,
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

func (s *sand) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider) model.Cell {
	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottom != nil && *bottom.CellType == model.EmptyCell.String() {
		return utils.CreateCellOnCellLocation(model.SandCell.String(), bottom)
	}

	bottomLeft := utils.GetBottomLeftNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottomLeft != nil && *bottomLeft.CellType == model.EmptyCell.String() {
		return utils.CreateCellOnCellLocation(model.SandCell.String(), bottomLeft)
	}

	bottomRight := utils.GetBottomRightNeighbour(currentGeneration, currentCell.X, currentCell.Y)
	if bottomRight != nil && *bottomRight.CellType == model.EmptyCell.String() {
		return utils.CreateCellOnCellLocation(model.SandCell.String(), bottomRight)
	}

	return utils.CreateCellOnCellLocation(model.SandCell.String(), &currentCell)
}
