package sort

func HeapSortInt(array []int) []int {
	var size int = len(array)
	// var i int = (size / 2) - 1
	var sorted []int = make([]int, size)
	copy(sorted, array)
	var max_idx int = 0
	for unsorted_length := size; unsorted_length > 1; unsorted_length-- {
		// build heap tree
		for i := (int(unsorted_length/2) - 1); i >= 0; i-- {
			max_idx = i
			if sorted[max_idx] < sorted[2*i+1] {
				max_idx = 2*i + 1
			}
			if (unsorted_length > 2*i+2) && sorted[max_idx] < sorted[2*i+2] {
				max_idx = 2*i + 2
			}

			if max_idx != i {
				sorted[i], sorted[max_idx] = sorted[max_idx], sorted[i]
			}
		}

		// swap root and end of slice
		sorted[0], sorted[unsorted_length-1] = sorted[unsorted_length-1], sorted[0]
	}
	return sorted
}
