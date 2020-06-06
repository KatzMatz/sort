package sort

import (
	"strings"
)

func SelectionSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold(ascOrDsc, "asc") {
		return SelectionSortIntAsc(slice)
	} else if strings.EqualFold(ascOrDsc, "dsc") {
		return SelectionSortIntDsc(slice)
	} else {
		return nil
	}
}

func SelectionSortIntAsc(slice []int) []int {
	var size int = len(slice)

	var minIdx int = 0
	for i := 0; i < size; i++ {
		minIdx = i
		for j := i + 1; j < size; j++ {
			if slice[minIdx] > slice[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}

	return slice
}

func SelectionSortIntDsc(slice []int) []int {
	var size int = len(slice)

	var maxIdx int = 0
	for i := 0; i < size; i++ {
		maxIdx = i
		for j := i + 1; j < size; j++ {
			if slice[maxIdx] < slice[j] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
		}
	}

	return slice
}

func SelectionSortFloat64(slice []float64, ascOrDsc string) []float64 {
	if strings.EqualFold(ascOrDsc, "asc") {
		return SelectionSortFloat64Asc(slice)
	} else if strings.EqualFold(ascOrDsc, "dsc") {
		return SelectionSortFloat64Dsc(slice)
	} else {
		return nil
	}
}

func SelectionSortFloat64Asc(slice []float64) []float64 {
	var size int = len(slice)

	var minIdx int = 0
	for i := 0; i < size; i++ {
		minIdx = i
		for j := i + 1; j < size; j++ {
			if slice[minIdx] > slice[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}

	return slice
}

func SelectionSortFloat64Dsc(slice []float64) []float64 {
	var size int = len(slice)

	var maxIdx int = 0
	for i := 0; i < size; i++ {
		maxIdx = i
		for j := i + 1; j < size; j++ {
			if slice[maxIdx] < slice[j] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
		}
	}

	return slice
}
