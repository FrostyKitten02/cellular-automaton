package model

type ElementProviderImpl struct {
	BurningElements            []Element
	BurningElementsCellTypes   []string
	FlammableElements          []Element
	FlammableElementsCellTypes []string
}

func (e *ElementProviderImpl) GetFlammableElements() []Element {
	return e.FlammableElements
}

func (e *ElementProviderImpl) GetFlammableElementosTypes() []string {
	return e.FlammableElementsCellTypes
}

func (e *ElementProviderImpl) IsFlammableCellType(cellType string) bool {
	for _, flammableCellType := range e.FlammableElementsCellTypes {
		if flammableCellType == cellType {
			return true
		}
	}

	return false
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
	burningElementTypes := elementsToCellType(burningElements)

	flammableElements := filterFlammableElements(elements)
	flammableElementTypes := elementsToCellType(flammableElements)

	return &ElementProviderImpl{
		BurningElements:            burningElements,
		BurningElementsCellTypes:   burningElementTypes,
		FlammableElements:          flammableElements,
		FlammableElementsCellTypes: flammableElementTypes,
	}
}

func elementsToCellType(elements []Element) []string {
	res := make([]string, len(elements))

	for i, element := range elements {
		res[i] = element.GetCellType().String()
	}

	return res
}

func filterBurningElements(elements []Element) []Element {
	return filterElements(elements, func(prop ElementProperties) bool {
		return prop.Burning
	})
}

func filterFlammableElements(elements []Element) []Element {
	return filterElements(elements, func(prop ElementProperties) bool {
		return prop.Flammable
	})
}

func filterElements(elements []Element, getProperty func(properties ElementProperties) bool) []Element {
	if elements == nil {
		return nil
	}

	res := make([]Element, 0)
	for _, element := range elements {
		config := element.GetProperties()
		if &config == nil {
			continue
		}

		if getProperty(element.GetProperties()) {
			res = append(res, element)
		}
	}

	return res
}
