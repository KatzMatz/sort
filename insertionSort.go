package sort

import (
	"strings"
)

func InsertionSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return InsertionSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return InsertionSortIntDsc(slice)
	} else {
		return nil
	}
}

func InsertionSortIntAsc(slice []int) []int {
	var size int = len(slice)

	var key int = 0
	for i := 1; i < size; i++ {
		key = i
		for (key > 0) && (slice[key] < slice[key-1]) {
			slice[key], slice[key-1] = slice[key-1], slice[key]
			key--
		}
	}
	return slice
}

func InsertionSortIntDsc(slice []int) []int {
	var size int = len(slice)

	var key int = 0
	for i := 1; i < size; i++ {
		key = i
		for (key > 0) && (slice[key] > slice[key-1]) {
			slice[key], slice[key-1] = slice[key-1], slice[key]
			key--
		}
	}
	return slice
}
