package sort

import (
	"strings"
)

// if ascOrDsc is correct, sort success and return true
// else sort fail and return false
func BubbleSortInt(slice []int, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		BubbleSortIntAsc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		BubbleSortIntDsc(slice)
		return true
	} else {
		return false
	}
}

func BubbleSortIntAsc(slice []int) {
	bubbleSortAsc(IntSlice(slice))
}

func BubbleSortIntDsc(slice []int) {
	bubbleSortDsc(IntSlice(slice))
}

func BubbleSortFloat64(slice []float64, ascOrDsc string) bool {
	if strings.EqualFold("asc", ascOrDsc) {
		BubbleSortFloat64Asc(slice)
		return true
	} else if strings.EqualFold("dsc", ascOrDsc) {
		BubbleSortFloat64Dsc(slice)
		return true
	} else {
		return false
	}
}

func BubbleSortFloat64Asc(slice []float64) {
	bubbleSortAsc(Float64Slice(slice))
}

func BubbleSortFloat64Dsc(slice []float64) {
	bubbleSortDsc(Float64Slice(slice))
}

func bubbleSortAsc(data SortInterface) {
	var n int = data.Len()

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func bubbleSortDsc(data SortInterface) {
	var n int = data.Len()

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}
