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

	result := BubbleSortIntAsc(sliceInt)

	for i, v := range result {
		if v != correctSliceIntAsc[i] {
			t.Fatal("failed")
		}
	}

}

func TestBubbleSortIntDsc(t *testing.T) {
	result := BubbleSortIntDsc(sliceInt)

	for i, v := range result {
		if v != correctSliceIntDsc[i] {
			t.Fatal("failed")
		}
	}
}

func TestBubbleSortInt(t *testing.T) {
	resultAsc := BubbleSortInt(sliceInt, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := BubbleSortInt(sliceInt, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if BubbleSortInt(sliceInt, "not asc or dsc") != nil {
		t.Fatal("failed")
	}
}

func TestBubbleSortFloat(t *testing.T) {
	resultAsc := BubbleSortFloat64(sliceFloat64, "asc")

	for i, v := range resultAsc {
		if v != correctSliceFloat64Asc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := BubbleSortFloat64(sliceFloat64, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceFloat64Dsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if BubbleSortFloat64(sliceFloat64, "not asc or dsc") != nil {
		t.Fatal("failed")
	}
}

func TestQuickSortInt(t *testing.T) {
	resultAsc := QuickSortInt(sliceInt, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := QuickSortInt(sliceInt, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if QuickSortInt(sliceInt, "not asc or dsc") != nil {
		t.Fatal("failed")
	}
}

func TestQuickSortFloat(t *testing.T) {
	resultAsc := QuickSortFloat64(sliceFloat64, "asc")

	for i, v := range resultAsc {
		if v != correctSliceFloat64Asc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := QuickSortFloat64(sliceFloat64, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceFloat64Dsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if QuickSortFloat64(sliceFloat64, "not asc or dsc") != nil {
		t.Fatal("failed")
	}
}

func TestReverseSliceInt(t *testing.T) {
	result := reverseSliceInt(correctSliceIntAsc)

	for i, v := range result {
		if v != correctSliceIntDsc[i] {
			t.Fatal("failed")
		}
	}
}

func TestHeapSortInt(t *testing.T) {
	resultAsc := HeapSortInt(sliceInt, "asc")

	for i, v := range resultAsc {
		if v != correctSliceIntAsc[i] {
			t.Errorf("%d\n", sliceInt)

			t.Errorf("%d\n", resultAsc)
			t.Errorf("%d\n", correctSliceIntAsc)

			t.Fatal("asc failed")
		}
	}

	resultDsc := HeapSortInt(sliceInt, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceIntDsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if HeapSortInt(sliceInt, "not asc or dsc") != nil {
		t.Fatal("failed")
	}
}

func TestHeapSortFloat64(t *testing.T) {
	resultAsc := HeapSortFloat64(sliceFloat64, "asc")

	for i, v := range resultAsc {
		if v != correctSliceFloat64Asc[i] {
			t.Fatal("asc failed")
		}
	}

	resultDsc := HeapSortFloat64(sliceFloat64, "dsc")

	for i, v := range resultDsc {
		if v != correctSliceFloat64Dsc[i] {
			t.Fatal("dsc failed")
		}
	}

	if HeapSortFloat64(sliceFloat64, "not asc or dsc") != nil {
		t.Fatal("Failed")
	}

}
