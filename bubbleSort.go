package sort

import (
	"strings"
)

func BubbleSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return BubbleSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return BubbleSortIntDsc(slice)
	} else {
		return nil
	}
}

func BubbleSortIntAsc(slice []int) []int {
	var n int = len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

func BubbleSortIntDsc(slice []int) []int {
	var n int = len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

func BubbleSortFloat64(slice []float64, ascOrDsc string) []float64 {
	if strings.EqualFold("asc", ascOrDsc) {
		return BubbleSortFloat64Asc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return BubbleSortFloat64Dsc(slice)
	} else {
		return nil
	}
}

func BubbleSortFloat64Asc(slice []float64) []float64 {
	var n int = len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

func BubbleSortFloat64Dsc(slice []float64) []float64 {
	var n int = len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}
