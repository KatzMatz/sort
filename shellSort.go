package sort

import (
	"strings"
)

func ShellSortInt(slice []int, ascOrDsc string) []int {
	if strings.EqualFold("asc", ascOrDsc) {
		return ShellSortIntAsc(slice)
	} else if strings.EqualFold("dsc", ascOrDsc) {
		return ShellSortIntDsc(slice)
	} else {
		return nil
	}
}

func ShellSortIntAsc(slice []int) []int {
	var size = len(slice)

	// decide h
	var h int = 1
	for h < (size / 9) {
		h = 3*h + 1
	}

	for ; h > 0; h = h / 3 {
		for i := 1; i < size; i++ {
			for key := i; (key >= h) && (slice[key] < slice[key-h]); key = key - h {
				slice[key], slice[key-h] = slice[key-h], slice[key]
			}
		}
	}

	return slice
}

func ShellSortIntDsc(slice []int) []int {
	var size = len(slice)

	// decide h
	var h int = 1
	for h < (size / 9) {
		h = 3*h + 1
	}

	for ; h > 0; h = h / 3 {
		for i := 1; i < size; i++ {
			for key := i; (key >= h) && (slice[key] > slice[key-h]); key = key - h {
				slice[key], slice[key-h] = slice[key-h], slice[key]
			}
		}
	}

	return slice
}
