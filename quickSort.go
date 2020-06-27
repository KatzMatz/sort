package sort

import (
	"strings"
)

func QuickSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		QuickSortAscInt(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		QuickSortDscInt(slice)
		return true
	} else {
		return false
	}
}

func QuickSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		QuickSortAscFloat64(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		QuickSortDscFloat64(slice)
		return true
	} else {
		return false
	}
}

func separateByPivotAsc(data SortInterface, startIdx int, endIdx int) int {
	var midIdx int = startIdx

	for i := startIdx + 1; i < endIdx; i++ {
		if data.Less(i, endIdx) {
			data.Swap(i, midIdx)
			midIdx++
		}
	}
	data.Swap(midIdx, endIdx)
	return midIdx
}

func separateByPivotDsc(data SortInterface, startIdx int, endIdx int) int {
	var midIdx int = startIdx

	for i := startIdx + 1; i < endIdx; i++ {
		if data.Less(endIdx, i) {
			data.Swap(i, midIdx)
			midIdx++
		}
	}
	data.Swap(midIdx, endIdx)
	return midIdx
}

func quickSortAscSub(data SortInterface, startIdx int, endIdx int) {
	if startIdx < endIdx {
		pivotIdx := separateByPivotAsc(data, startIdx, endIdx)
		quickSortAscSub(data, startIdx, pivotIdx)
		quickSortAscSub(data, pivotIdx+1, endIdx)
	}
}

func quickSortDscSub(data SortInterface, startIdx int, endIdx int) {
	if startIdx < endIdx {
		pivotIdx := separateByPivotDsc(data, startIdx, endIdx)
		quickSortDscSub(data, startIdx, pivotIdx)
		quickSortDscSub(data, pivotIdx+1, endIdx)
	}
}

func quickSortAsc(data SortInterface) {
	quickSortAscSub(data, 0, data.Len()-1)
}

func quickSortDsc(data SortInterface) {
	quickSortDscSub(data, 0, data.Len()-1)
}

func QuickSortAscInt(slice []int) {
	quickSortAsc(IntSlice(slice))
}

func QuickSortDscInt(slice []int) {
	quickSortDsc(IntSlice(slice))
}

func QuickSortAscFloat64(slice []float64) {
	quickSortAsc(Float64Slice(slice))
}

func QuickSortDscFloat64(slice []float64) {
	quickSortDsc(Float64Slice(slice))
}
