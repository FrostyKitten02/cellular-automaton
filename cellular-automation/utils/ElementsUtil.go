package utils

import "cellular-automation/model"

func FindElementForCellType(elements *[]model.Element, cellType string) *model.Element {
	for _, element := range *elements {
		if cellType != element.GetCellType().String() {
			continue
		}

		return &element
	}

	return nil
}
