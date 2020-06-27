package sort

import (
	"strings"
)

func shellSortAsc(data SortInterface) {
	var size = data.Len()

	var h int = 1
	for h < (size / 9) {
		h = 3*h + 1
	}

	for ; h > 0; h = h / 3 {
		for i := 1; i < size; i++ {
			for key := i; (key >= h) && data.Less(key, key-h); key = key - h {
				data.Swap(key, key-h)
			}
		}
	}
}

func shellSortDsc(data SortInterface) {
	var size = data.Len()

	var h int = 1
	for h < (size / 9) {
		h = 3*h + 1
	}

	for ; h > 0; h = h / 3 {
		for i := 1; i < size; i++ {
			for key := i; (key >= h) && data.Less(key-h, key); key = key - h {
				data.Swap(key, key-h)
			}
		}
	}
}

func ShellSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		ShellSortAscInt(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		ShellSortDscInt(slice)
		return true
	} else {
		return false
	}
}

func ShellSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		ShellSortAscFloat64(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		ShellSortDscFloat64(slice)
		return true
	} else {
		return false
	}
}

func ShellSortAscInt(slice []int) {
	shellSortAsc(IntSlice(slice))
}

func ShellSortDscInt(slice []int) {
	shellSortDsc(IntSlice(slice))
}

func ShellSortAscFloat64(slice []float64) {
	shellSortAsc(Float64Slice(slice))
}

func ShellSortDscFloat64(slice []float64) {
	shellSortDsc(Float64Slice(slice))
}
