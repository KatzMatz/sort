# Golang sort package

- bubble sort
- selection sort
- insertion sort
- shell sort
- merge sort
- heap sort
- quick sort

## Usage

```go
var slice = []int{2,3,4,1}
slice = sort.BubbleSortIntAsc(slice) // => [1, 2, 3, 4]
// or
// slice = sort.bubbleSortInt(slice, "asc")
```

```go
var slice = []float64{2.1, 3.1, 4.4, 1.4}
slice = sort.QuickSortFloat64Dsc(slice) // => [4.4, 3.1, 2.1, 1.4]
```