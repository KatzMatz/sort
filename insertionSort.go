package sort

import (
	"strings"
)

func InsertionSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		InsertionSortIntAsc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		InsertionSortIntDsc(slice)
		return true
	} else {
		return false
	}
}

func InsertionSortIntAsc(slice []int) {
	InsertionSortAsc(IntSlice(slice))
}

func InsertionSortIntDsc(slice []int) {
	InsertionSortDsc(IntSlice(slice))
}

func InsertionSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		InsertionSortFloat64Asc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		InsertionSortFloat64Dsc(slice)
		return true
	} else {
		return false
	}
}

func InsertionSortFloat64Asc(slice []float64) {
	InsertionSortAsc(Float64Slice(slice))
}

func InsertionSortFloat64Dsc(slice []float64) {
	InsertionSortDsc(Float64Slice(slice))
}

func InsertionSortAsc(data SortInterface) {
	var size int = data.Len()

	var key int = 0

	for i := 1; i < size; i++ {
		key = i
		for (key > 0) && (data.Less(key, key-1)) {
			data.Swap(key, key-1)
			key--
		}
	}
}

func InsertionSortDsc(data SortInterface) {
	var size int = data.Len()

	var key int = 0

	for i := 1; i < size; i++ {
		key = i
		for (key > 0) && (data.Less(key-1, key)) {
			data.Swap(key, key-1)
			key--
		}
	}
}
