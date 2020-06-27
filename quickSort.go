package sort

import (
	"strings"
)

func QuickSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return QuickSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return QuickSortIntDsc(slice)
	} else {
		return nil
	}
}

func QuickSortIntAsc(slice []int) []int {
	var size int = len(slice)
	if size > 1 {
		var sorted []int = make([]int, size)

		var pivot int = slice[0]
		var left []int
		var right []int

		for _, item := range slice[1:] {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortIntAsc(left)
		right = QuickSortIntAsc(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return slice
	}

}

func QuickSortIntDsc(slice []int) []int {
	var size int = len(slice)
	if size > 1 {
		var sorted []int = make([]int, size)

		var pivot int = slice[0]
		var left []int
		var right []int

		for _, item := range slice[1:] {
			if item >= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortIntDsc(left)
		right = QuickSortIntDsc(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return slice
	}

}

func QuickSortFloat64(slice []float64, ascOrDsc string) []float64 {
	if strings.EqualFold("asc", ascOrDsc) {
		return QuickSortFloat64Asc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return QuickSortFloat64Dsc(slice)
	} else {
		return nil
	}
}

func QuickSortFloat64Asc(slice []float64) []float64 {
	var size int = len(slice)
	if size > 1 {
		var sorted []float64 = make([]float64, size)

		var pivot float64 = slice[0]
		var left []float64
		var right []float64

		for _, item := range slice[1:] {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortFloat64Asc(left)
		right = QuickSortFloat64Asc(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return slice
	}

}

func QuickSortFloat64Dsc(slice []float64) []float64 {
	var size int = len(slice)
	if size > 1 {
		var sorted []float64 = make([]float64, size)

		var pivot float64 = slice[0]
		var left []float64
		var right []float64

		for _, item := range slice[1:] {
			if item >= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortFloat64Dsc(left)
		right = QuickSortFloat64Dsc(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return slice
	}

}

func separatePivot(data SortInterface, startIdx int, endIdx int) int {

}

func QuickSortAsc(data SortInterface, startIdx, endIdx int) {
	var size int = endIdx - startIdx
	if size > 1 {
		// var sorted []float64 = make([]float64, size)

		var pivot float64 = slice[0]
		var left []float64
		var right []float64

		for _, item := range slice[1:] {
			if item <= pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = QuickSortFloat64Asc(left)
		right = QuickSortFloat64Asc(right)

		sorted = append(left, pivot)
		sorted = append(sorted, right...)

		return sorted
	} else {
		return slice
	}
}
