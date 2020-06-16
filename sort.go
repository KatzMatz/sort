package sort

type SortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// IntSlice implements SortInterface
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Float64Slice implements SortInterface
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reverseSlice(data SortInterface) {
	var size int = data.Len()
	for i := 0; i < size/2; i++ {
		data.Swap(i, size-i-1)
	}
}

func reverseSliceInt(slice []int) {
	reverseSlice(IntSlice(slice))
}

func reverseSliceFloat64(slice []float64) {
	reverseSlice(Float64Slice(slice))
}
