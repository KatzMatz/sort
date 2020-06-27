package sort

import (
	"strings"
)

func MergeSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return MergeSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return MergeSortIntDsc(slice)
	} else {
		return nil
	}
}

func MergeSortIntAsc(slice []int) []int {
	var size int = len(slice)

	if size > 1 {
		// Split part
		var midIdx int = size / 2
		var leftSlice []int = MergeSortIntAsc(slice[:midIdx])
		var rightSlice []int = MergeSortIntAsc(slice[midIdx:])

		// Merge part
		var sorted []int = make([]int, size)
		var leftIdx int = 0
		var rightIdx int = 0

		for i := 0; i < size; i++ {
			if (rightIdx == size-midIdx) || (leftIdx < midIdx && leftSlice[leftIdx] <= rightSlice[rightIdx]) {
				sorted[i] = leftSlice[leftIdx]
				leftIdx++
			} else {
				sorted[i] = rightSlice[rightIdx]
				rightIdx++
			}
		}
		return sorted

	} else {
		return slice
	}

}

func MergeSortIntDsc(slice []int) []int {
	var size int = len(slice)

	if size > 1 {
		// Split part
		var midIdx int = size / 2
		var leftSlice []int = MergeSortIntDsc(slice[:midIdx])
		var rightSlice []int = MergeSortIntDsc(slice[midIdx:])

		// Merge part
		var sorted []int = make([]int, size)
		var leftIdx int = 0
		var rightIdx int = 0

		for i := 0; i < size; i++ {
			if (rightIdx == size-midIdx) || (leftIdx < midIdx && leftSlice[leftIdx] >= rightSlice[rightIdx]) {
				sorted[i] = leftSlice[leftIdx]
				leftIdx++
			} else {
				sorted[i] = rightSlice[rightIdx]
				rightIdx++
			}
		}
		return sorted

	} else {
		return slice
	}

}

func MergeSortFloat64(slice []float64, ascOrDsc string) []float64 {
	if strings.EqualFold("asc", ascOrDsc) {
		return MergeSortFloat64Asc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return MergeSortFloat64Dsc(slice)
	} else {
		return nil
	}
}

func MergeSortFloat64Asc(slice []float64) []float64 {
	var size int = len(slice)

	if size > 1 {
		// Split part
		var midIdx int = size / 2
		var leftSlice []float64 = MergeSortFloat64Asc(slice[:midIdx])
		var rightSlice []float64 = MergeSortFloat64Asc(slice[midIdx:])

		// Merge part
		var sorted []float64 = make([]float64, size)
		var leftIdx int = 0
		var rightIdx int = 0

		for i := 0; i < size; i++ {
			if (rightIdx == size-midIdx) || (leftIdx < midIdx && leftSlice[leftIdx] <= rightSlice[rightIdx]) {
				sorted[i] = leftSlice[leftIdx]
				leftIdx++
			} else {
				sorted[i] = rightSlice[rightIdx]
				rightIdx++
			}
		}
		return sorted

	} else {
		return slice
	}

}

func MergeSortFloat64Dsc(slice []float64) []float64 {
	var size int = len(slice)

	if size > 1 {
		// Split part
		var midIdx int = size / 2
		var leftSlice []float64 = MergeSortFloat64Dsc(slice[:midIdx])
		var rightSlice []float64 = MergeSortFloat64Dsc(slice[midIdx:])

		// Merge part
		var sorted []float64 = make([]float64, size)
		var leftIdx int = 0
		var rightIdx int = 0

		for i := 0; i < size; i++ {
			if (rightIdx == size-midIdx) || (leftIdx < midIdx && leftSlice[leftIdx] >= rightSlice[rightIdx]) {
				sorted[i] = leftSlice[leftIdx]
				leftIdx++
			} else {
				sorted[i] = rightSlice[rightIdx]
				rightIdx++
			}
		}
		return sorted

	} else {
		return slice
	}

}

func MergeSortAsc(data SortInterface, startIdx int, endIdx int, dataSize int) {
	var size int = endIdx - startIdx

	if size > 1 {
		// Split part
		var midIdx int = size / 2
		MergeSortAsc(data, 0, midIdx, midIdx)
		MergeSortAsc(data, midIdx, size, size-midIdx)

		// Merge part
		var leftIdx int = 0
		var rightIdx int = 0

		for i := 0; i < size; i++ {
			if (rightIdx == size-midIdx) || (leftIdx < midIdx && leftSlice[leftIdx] <= rightSlice[rightIdx]) {
				sorted[i] = leftSlice[leftIdx]
				leftIdx++
			} else {
				sorted[i] = rightSlice[rightIdx]
				rightIdx++
			}
		}

	}
}
