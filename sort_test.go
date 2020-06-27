package sort

import (
	"testing"
)

var sliceInt []int = []int{3, 2, 1, 5, 4, 9, 8, 7, 6}
var correctSliceIntAsc []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var correctSliceIntDsc []int = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

var sliceFloat64 []float64 = []float64{3.0, 2.0, 1.0, 5.0, 4.0, 9.0, 8.0, 7.0, 6.0}
var correctSliceFloat64Asc []float64 = []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
var correctSliceFloat64Dsc []float64 = []float64{9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0, 1.0}

func TestBubbleSortIntAsc(t *testing.T) {

	result := make([]int, len(sliceInt))
	copy(result, sliceInt)

	BubbleSortIntAsc(result)

	for i, v := range result {
		if v != correctSliceIntAsc[i] {
			t.Fatal("failed")
		}
	}

}

func TestBubbleSortIntDsc(t *testing.T) {
	result := make([]int, len(sliceInt))
	copy(result, sliceInt)

	BubbleSortIntDsc(result)

	for i, v := range result {
		if v != correctSliceIntDsc[i] {
			t.Fatal("failed")
		}
	}
}

func TestBubbleSortInt(t *testing.T) {
	resultAsc := make([]int, len(sliceInt))
	copy(resultAsc, sliceInt)

	BubbleSortInt(resultAsc, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]int, len(sliceInt))
	copy(resultDsc, sliceInt)
	BubbleSortInt(resultDsc, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if BubbleSortInt(sliceInt, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}

func TestBubbleSortFloat(t *testing.T) {
	resultAsc := make([]float64, len(sliceFloat64))
	copy(resultAsc, sliceFloat64)

	BubbleSortFloat64(resultAsc, "asc")

	for i, v := range resultAsc {
		if v != correctSliceFloat64Asc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]float64, len(sliceFloat64))
	copy(resultDsc, sliceFloat64)

	BubbleSortFloat64(resultDsc, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceFloat64Dsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if BubbleSortFloat64(sliceFloat64, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}

func TestQuickSortInt(t *testing.T) {
	resultAsc := make([]int, len(sliceInt))
	copy(resultAsc, sliceInt)

	QuickSortInt(resultAsc, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Errorf("%d\n", sliceInt)
			t.Errorf("%d\n", correctSliceIntAsc)
			t.Errorf("%d\n", resultAsc)
			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]int, len(sliceInt))
	copy(resultDsc, sliceInt)

	QuickSortInt(resultDsc, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if QuickSortInt(sliceInt, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}

func TestQuickSortFloat(t *testing.T) {
	resultAsc := make([]float64, len(sliceFloat64))
	copy(resultAsc, sliceFloat64)

	QuickSortFloat64(resultAsc, "asc")

	for i, v := range resultAsc {
		if v != correctSliceFloat64Asc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]float64, len(sliceFloat64))
	copy(resultDsc, sliceFloat64)

	QuickSortFloat64(resultDsc, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceFloat64Dsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if QuickSortFloat64(sliceFloat64, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}

func TestReverseSliceInt(t *testing.T) {
	result := make([]int, len(sliceInt))
	copy(result, sliceInt)

	reverseSliceInt(result)

	for i, v := range result {
		if v != sliceInt[len(sliceInt)-1-i] {
			t.Errorf("%d\n", sliceInt)
			t.Errorf("%d\n", result)
			t.Fatal("failed")
		}
	}
}

func TestHeapSortInt(t *testing.T) {
	resultAsc := make([]int, len(sliceInt))
	copy(resultAsc, sliceInt)
	HeapSortInt(resultAsc, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Errorf("%d\n", sliceInt)

			t.Errorf("%d\n", resultAsc)
			t.Errorf("%d\n", correctSliceIntAsc)

			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]int, len(sliceInt))
	copy(resultDsc, sliceInt)
	HeapSortInt(resultDsc, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if HeapSortInt(sliceInt, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}

func TestShellSortInt(t *testing.T) {
	resultAsc := make([]int, len(sliceInt))
	copy(resultAsc, sliceInt)

	ShellSortAscInt(resultAsc)

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := make([]int, len(sliceInt))
	copy(resultDsc, sliceInt)

	ShellSortDscInt(resultDsc)

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if ShellSortInt(resultDsc, "not asc or dsc") != false {
		t.Fatal("failed")
	}

}

func TestSelectionSortInt(t *testing.T) {
	resultAsc := make([]int, len(sliceInt))
	copy(resultAsc, sliceInt)

	SelectionSortAscInt(resultAsc)

	if isSortedAscInt(resultAsc) == false {
		t.Fatal("asc failed")
	}

	resultDsc := make([]int, len(sliceInt))
	copy(resultDsc, sliceInt)

	SelectionSortDscInt(resultDsc)

	if isSortedDscInt(resultDsc) == false {
		t.Fatal("dsc failed")
	}

	if SelectionSortInt(resultAsc, "not asc or dsc") != false {
		t.Fatal("failed")
	}
}
