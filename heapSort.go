package sort

import (
	"strings"
)

func HeapSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return HeapSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return reverseSliceInt(HeapSortIntAsc(slice))
	} else {
		return nil
	}
}

func HeapSortIntAsc(slice []int) []int {
	var size int = len(slice)
	var maxIdx int = 0

	for unsortedLength := size; unsortedLength > 1; unsortedLength-- {
		// build heap tree
		for i := (int(unsortedLength/2) - 1); i >= 0; i-- {
			maxIdx = i
			if slice[maxIdx] < slice[2*i+1] {
				maxIdx = 2*i + 1
			}
			if (unsortedLength > 2*i+2) && slice[maxIdx] < slice[2*i+2] {
				maxIdx = 2*i + 2
			}

			if maxIdx != i {
				slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
			}
		}

		// swap root and end of slice
		slice[0], slice[unsortedLength-1] = slice[unsortedLength-1], slice[0]
	}
	return slice
}

func reverseSliceInt(slice []int) []int {
	var size int = len(slice)
	for i := 0; i < size/2; i++ {
		slice[i], slice[size-i] = slice[size-i], slice[i]
	}
	return slice
}

func HeapSortFloat64(slice []float64, ascOrDsc string) []float64 {
	if strings.EqualFold("asc", ascOrDsc) {
		return HeapSortFloat64Asc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return reverseSliceFloat64(HeapSortFloat64Asc(slice))
	} else {
		return nil
	}
}

func HeapSortFloat64Asc(slice []float64) []float64 {
	var size int = len(slice)
	var maxIdx int = 0

	for unsortedLength := size; unsortedLength > 1; unsortedLength-- {
		// build heap tree
		for i := (int(unsortedLength/2) - 1); i >= 0; i-- {
			maxIdx = i
			if slice[maxIdx] < slice[2*i+1] {
				maxIdx = 2*i + 1
			}
			if (unsortedLength > 2*i+2) && slice[maxIdx] < slice[2*i+2] {
				maxIdx = 2*i + 2
			}

			if maxIdx != i {
				slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
			}
		}

		// swap root and end of slice
		slice[0], slice[unsortedLength-1] = slice[unsortedLength-1], slice[0]
	}
	return slice
}

func reverseSliceFloat64(slice []float64) []float64 {
	var size int = len(slice)
	for i := 0; i < size/2; i++ {
		slice[i], slice[size-i] = slice[size-i], slice[i]
	}
	return slice
}
