package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"math/rand"
)

var Fire = fire{
	cellType: model.FireCell,
	properties: model.ElementProperties{
		Flammable: false,
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

func (f *fire) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo) model.Cell {
	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	//replace with flame if hits!!
	if bottom != nil {
		if provider.IsFlammableCellType(*bottom.CellType) {
			return utils.CreateCellOnCellLocation(model.WhiteSmoke.String(), bottom, gameInfo.GenerationNum)
		}

		return utils.CreateCellOnCellLocation(model.DarkSmoke.String(), bottom, gameInfo.GenerationNum)
	}

	possibleMoves := make([]model.Cell, 0)
	if bottom != nil && *bottom.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *bottom)
	}

	leftBottom := utils.GetBottomLeftNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if leftBottom != nil && *leftBottom.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *leftBottom)
	}

	rightBottom := utils.GetBottomRightNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if rightBottom != nil && *rightBottom.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *rightBottom)
	}

	possibleMovesCount := len(possibleMoves)
	if possibleMovesCount == 0 {
		return utils.CreateCellOnCellLocation(f.GetCellType().String(), &currentCell, currentCell.BornGeneration)
	}

	randIndex := rand.Intn(possibleMovesCount)
	moveTo := possibleMoves[randIndex]
	return utils.CreateCellOnCellLocation(f.GetCellType().String(), &moveTo, currentCell.BornGeneration)
}
