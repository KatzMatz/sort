package sort

import (
	"strings"
)

func selectionSortAsc(data SortInterface) {
	var size int = data.Len()

	var minIdx int = 0
	for i := 0; i < size; i++ {
		minIdx = i
		for j := i + 1; j < size; j++ {
			if data.Less(j, minIdx) {
				minIdx = j
			}
		}
		if minIdx != i {
			data.Swap(i, minIdx)
		}
	}
}

func selectionSortDsc(data SortInterface) {
	var size int = data.Len()

	var maxIdx int = 0
	for i := 0; i < size; i++ {
		maxIdx = i
		for j := i + 1; j < size; j++ {
			if data.Less(maxIdx, j) {
				maxIdx = j
			}
		}
		if maxIdx != i {
			data.Swap(i, maxIdx)
		}
	}
}

func SelectionSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold(ascOrDsc, "asc") {
		SelectionSortAscInt(slice)
		return true
	} else if strings.EqualFold(ascOrDsc, "dsc") {
		SelectionSortDscInt(slice)
		return true
	} else {
		return false
	}
}

func SelectionSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold(ascOrDsc, "asc") {
		SelectionSortAscFloat64(slice)
		return true
	} else if strings.EqualFold(ascOrDsc, "dsc") {
		SelecitonSortDscFloat64(slice)
		return true
	} else {
		return false
	}
}

func SelectionSortAscInt(slice []int) {
	selectionSortAsc(IntSlice(slice))
}

func SelectionSortDscInt(slice []int) {
	selectionSortDsc(IntSlice(slice))
}

func SelectionSortAscFloat64(slice []float64) {
	selectionSortAsc(Float64Slice(slice))
}

func SelecitonSortDscFloat64(slice []float64) {
	selectionSortDsc(Float64Slice(slice))
}
