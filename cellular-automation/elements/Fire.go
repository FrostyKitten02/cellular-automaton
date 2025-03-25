package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"math/rand"
)

var Fire = fire{
	cellType: model.FireCell,
	properties: model.ElementProperties{
		Flameable: false,
		Burning:   true,
	},
}

type fire struct {
	cellType   model.CellType
	properties model.ElementProperties
}

func (f *fire) GetProperties() model.ElementProperties {
	return f.properties
}

func (f *fire) GetCellType() model.CellType {
	return f.cellType
}

func (f *fire) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider) model.Cell {
	moveIndex := rand.Intn(4)
	leftBottom := utils.GetBottomLeftNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if leftBottom != nil && *leftBottom.CellType == model.EmptyCell.String() && moveIndex <= 0 {
		return utils.CreateCellOnCellLocation(f.GetCellType().String(), leftBottom)
	}

	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if bottom != nil && *bottom.CellType == model.EmptyCell.String() && moveIndex <= 1 {
		return utils.CreateCellOnCellLocation(f.GetCellType().String(), bottom)
	}

	rightBottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if rightBottom != nil && *rightBottom.CellType == model.EmptyCell.String() && moveIndex <= 2 {
		return utils.CreateCellOnCellLocation(f.GetCellType().String(), rightBottom)
	}

	return utils.CreateCellOnCellLocation(f.GetCellType().String(), &currentCell)
}
