package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
)

var Wood = wood{
	cellType: model.WoodCell,
}

type wood struct {
	cellType model.CellType
}

func (w *wood) GetCellType() model.CellType {
	return w.cellType
}

func (w *wood) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell) model.Cell {
	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if bottom != nil && *bottom.CellType == model.EmptyCell.String() {
		return utils.CreateCellOnCellLocation(w.GetCellType().String(), bottom)
	}

	return utils.CreateCellOnCellLocation(w.GetCellType().String(), &currentCell)
}
