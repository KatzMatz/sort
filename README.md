# Golang sort package

- bubble sort
- selection sort
- insertion sort
- shell sort
- heap sort
- quick sort

## Usage

```go
var slice = []int{2,3,4,1}
sort.BubbleSortAscInt(slice) // => [1, 2, 3, 4]
// or
// sort.bubbleSortInt(slice, "asc")
```

```go
var slice = []float64{2.1, 3.1, 4.4, 1.4}
sort.QuickSortDscFloat64(slice) // => [4.4, 3.1, 2.1, 1.4]
```
