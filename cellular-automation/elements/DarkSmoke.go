package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"math/rand"
)

var DarkSmoke = darkSmoke{
	cellType: model.DarkSmoke,
	properties: model.ElementProperties{
		Flammable: true,
		Burning:   false,
	},
}

type darkSmoke struct {
	cellType   model.CellType
	properties model.ElementProperties
}

func (s *darkSmoke) GetProperties() model.ElementProperties {
	return s.properties
}

func (s *darkSmoke) GetCellType() model.CellType {
	return s.cellType
}

func (s *darkSmoke) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo) model.Cell {
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
			return utils.CreateCellOnCellLocation(s.GetCellType().String(), &currentCell, currentCell.BornGeneration)
		}

		randIndex := rand.Intn(possibleSideMovesCount)
		moveTo := possibleSideMoves[randIndex]
		return utils.CreateCellOnCellLocation(s.GetCellType().String(), &moveTo, currentCell.BornGeneration)
	}

	randIndex := rand.Intn(possibleMovesCount)
	moveTo := possibleMoves[randIndex]
	return utils.CreateCellOnCellLocation(s.GetCellType().String(), &moveTo, currentCell.BornGeneration)
}
