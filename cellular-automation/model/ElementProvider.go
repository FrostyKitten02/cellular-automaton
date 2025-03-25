package model

type ElementProviderImpl struct {
	BurningElements          []Element
	BurningElementsCellTypes []string
}

func (e *ElementProviderImpl) GetBurningElementsCellTypes() []string {
	return e.BurningElementsCellTypes
}

func (e *ElementProviderImpl) GetBurningElements() []Element {
	return e.BurningElements
}

func (e *ElementProviderImpl) IsBurningCellType(cellType string) bool {
	for _, burningCellType := range e.BurningElementsCellTypes {
		if burningCellType == cellType {
			return true
		}
	}

	return false
}

func NewElementProvider(elements []Element) ElementProvider {
	burningElements := filterBurningElements(elements)
	burningElementTypes := burningElementsToCellTypes(burningElements)

	return &ElementProviderImpl{
		BurningElements:          burningElements,
		BurningElementsCellTypes: burningElementTypes,
	}
}

func burningElementsToCellTypes(elements []Element) []string {
	res := make([]string, len(elements))

	for i, element := range elements {
		res[i] = element.GetCellType().String()
	}

	return res
}

func filterBurningElements(elements []Element) []Element {
	if elements == nil {
		return nil
	}

	res := make([]Element, 0)
	for _, element := range elements {
		config := element.GetProperties()
		if &config == nil {
			continue
		}

		if element.GetProperties().Burning {
			res = append(res, element)
		}
	}

	return res
}
