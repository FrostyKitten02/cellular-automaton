package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
)

var Wood = wood{
	cellType: model.WoodCell,
	properties: model.ElementProperties{
		Flammable: true,
		Burning:   false,
	},
}

type wood struct {
	cellType   model.CellType
	properties model.ElementProperties
}

func (w *wood) GetProperties() model.ElementProperties {
	return w.properties
}

func (w *wood) GetCellType() model.CellType {
	return w.cellType
}

func (w *wood) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo, futureGen *[][]model.Cell) {
	shouldBurn := utils.AnyBurningNeighbours(currentGeneration, currentCell, provider)

	if shouldBurn {
		sameLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
		utils.AppendCellInArr(&sameLocation, nil, futureGen)
		return
	}

	x := currentCell.GetX()
	y := currentCell.GetY()
	bottom := utils.GetBottomNeighbour(currentGeneration, x, y)
	if bottom != nil && *bottom.CellType == model.EmptyCell.String() {
		newLocation := utils.CreateCellOnCellLocation(w.GetCellType().String(), bottom, currentCell.BornGeneration)
		oldLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
		utils.AppendCellInArr(&newLocation, &oldLocation, futureGen)
		return
	}

	oldLocation := utils.CreateCellOnCellLocation(w.GetCellType().String(), &currentCell, currentCell.BornGeneration)
	utils.AppendCellInArr(&oldLocation, nil, futureGen)
}
