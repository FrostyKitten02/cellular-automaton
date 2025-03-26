package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"math/rand"
)

var WhiteSmoke = darkSmoke{
	cellType: model.WhiteSmoke,
	properties: model.ElementProperties{
		Flammable: true,
		Burning:   false,
	},
}

type whiteSmoke struct {
	cellType   model.CellType
	properties model.ElementProperties
}

func (s *whiteSmoke) GetProperties() model.ElementProperties {
	return s.properties
}

func (s *whiteSmoke) GetCellType() model.CellType {
	return s.cellType
}

func (s *whiteSmoke) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo, futureGen *[][]model.Cell) {
	possibleMoves := make([]model.Cell, 0)
	topLeft := utils.GetTopLeftNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if topLeft != nil && *topLeft.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *topLeft)
	}

	top := utils.GetTopNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if top != nil && *top.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *top)
	}

	topRight := utils.GetTopRightNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	if topRight != nil && *topRight.CellType == model.EmptyCell.String() {
		possibleMoves = append(possibleMoves, *topRight)
	}

	possibleMovesCount := len(possibleMoves)
	//if it cannot move up we check if we can move sideways randomly ofc!
	if possibleMovesCount == 0 {
		possibleSideMoves := make([]model.Cell, 0)

		left := utils.GetLeftNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
		if left != nil && *left.CellType == model.EmptyCell.String() {
			possibleSideMoves = append(possibleSideMoves, *left)
		}

		right := utils.GetRightNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
		if right != nil && *right.CellType == model.EmptyCell.String() {
			possibleSideMoves = append(possibleSideMoves, *right)
		}

		possibleSideMovesCount := len(possibleSideMoves)
		if possibleSideMovesCount == 0 {
			newLocation := utils.CreateCellOnCellLocation(s.GetCellType().String(), &currentCell, currentCell.BornGeneration)
			oldLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
			utils.AppendCellInArr(&newLocation, &oldLocation, futureGen)
			return
		}

		randIndex := rand.Intn(possibleSideMovesCount)
		moveTo := possibleSideMoves[randIndex]

		newLocation := utils.CreateCellOnCellLocation(s.GetCellType().String(), &moveTo, currentCell.BornGeneration)
		oldLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
		utils.AppendCellInArr(&newLocation, &oldLocation, futureGen)
		return
	}

	randIndex := rand.Intn(possibleMovesCount)
	moveTo := possibleMoves[randIndex]
	newLocation := utils.CreateCellOnCellLocation(s.GetCellType().String(), &moveTo, currentCell.BornGeneration)
	oldLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
	utils.AppendCellInArr(&newLocation, &oldLocation, futureGen)
}
