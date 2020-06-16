package sort

import (
	"strings"
)

func HeapSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		HeapSortIntAsc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		HeapSortIntDsc(slice)
		return true
	} else {
		return false
	}
}

func HeapSortIntAsc(slice []int) {
	HeapSortAsc(IntSlice(slice))
}

func HeapSortIntDsc(slice []int) {
	HeapSortAsc(IntSlice(slice))
	reverseSliceInt(slice)
}

func HeapSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		HeapSortFloat64Asc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		HeapSortFloat64Dsc(slice)
		return true
	} else {
		return false
	}
}

func HeapSortFloat64Asc(slice []float64) {
	HeapSortAsc(Float64Slice(slice))
}

func HeapSortFloat64Dsc(slice []float64) {
	HeapSortAsc(Float64Slice(slice))
	reverseSliceFloat64(slice)
}

func HeapSortAsc(data SortInterface) {
	var size int = data.Len()
	var maxIdx int = 0

	for unsortedLength := size; unsortedLength > 1; unsortedLength-- {
		// build heap tree
		for i := (int(unsortedLength/2) - 1); i >= 0; i-- {
			maxIdx = i
			if data.Less(maxIdx, 2*i+1) {
				maxIdx = 2*i + 1
			}
			if (unsortedLength > 2*i+2) && data.Less(maxIdx, 2*i+2) {
				maxIdx = 2*i + 2
			}

			if maxIdx != i {
				data.Swap(i, maxIdx)
			}
		}

		// swap root and end of slice
		data.Swap(0, unsortedLength-1)
	}
}
