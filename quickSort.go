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
