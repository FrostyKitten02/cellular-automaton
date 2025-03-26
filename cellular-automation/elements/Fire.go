package elements

import (
	"cellular-automation/model"
	"cellular-automation/utils"
	"log"
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

func (f *fire) NextGenerationCell(currentGeneration model.Grid, currentCell model.Cell, provider model.ElementProvider, gameInfo model.GameInfo, futureGen *[][]model.Cell) {
	bottom := utils.GetBottomNeighbour(currentGeneration, currentCell.GetX(), currentCell.GetY())
	//replace with flame if hits!!
	log.Println("Fire logic")
	if bottom != nil && *bottom.CellType != model.EmptyCell.String() {
		log.Print("Fire converting!")
		if provider.IsFlammableCellType(*bottom.CellType) {
			newLocation := utils.CreateCellOnCellLocation(model.WhiteSmoke.String(), &currentCell, gameInfo.GenerationNum)
			utils.AppendCellInArr(&newLocation, nil, futureGen)
		}

		newLocation := utils.CreateCellOnCellLocation(model.DarkSmoke.String(), &currentCell, gameInfo.GenerationNum)
		utils.AppendCellInArr(&newLocation, nil, futureGen)
		return
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
		//THIS SHOULDN'T HAPPEN
		same := utils.CreateCellOnCellLocation(f.GetCellType().String(), &currentCell, currentCell.BornGeneration)
		utils.AppendCellInArr(&same, nil, futureGen)
		return
	}

	randIndex := rand.Intn(possibleMovesCount)
	moveTo := possibleMoves[randIndex]
	newLocation := utils.CreateCellOnCellLocation(f.GetCellType().String(), &moveTo, currentCell.BornGeneration)
	oldLocation := utils.CreateCellOnCellLocation(model.EmptyCell.String(), &currentCell, gameInfo.GenerationNum)
	utils.AppendCellInArr(&newLocation, &oldLocation, futureGen)
}
