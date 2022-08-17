package findmediansortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays1(t *testing.T) {
	assert.Equal(t, float64(11), findMedianSortedArrays([]int{7, 9, 11}, []int{1, 3, 5, 11, 13, 15, 17, 19, 21}))
}

func Test_findMedianSortedArrays2(t *testing.T) {
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func Test_findMedianSortedArrays3(t *testing.T) {
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
func Test_findMedianSortedArrays4(t *testing.T) {
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{1}, []int{}))
}
func Test_findMedianSortedArrays5(t *testing.T) {
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{}, []int{1}))
}

func Test_findMedianSortedArrays6(t *testing.T) {
	assert.Equal(t, float64(6), findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{6}))
}

func Test_findMedianSortedArrays7(t *testing.T) {
	assert.Equal(t, float64(7), findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{1000}))
}

func Test_findMedianSortedArrays8(t *testing.T) {
	assert.Equal(t, float64(6.5), findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{2, 4, 6, 8, 10, 12}))
}

func Test_findMedianSortedArrays9(t *testing.T) {
	assert.Equal(t, float64(6), findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{1, 3, 5, 7, 9, 11}))

}
func Test_findMedianSortedArrays10(t *testing.T) {
	assert.Equal(t, float64(11), findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{11, 13, 15, 17, 19, 21}))
}

func Test_findMedianSortedArrays11(t *testing.T) {
	assert.Equal(t, float64(11), findMedianSortedArrays([]int{7, 9, 11}, []int{1, 3, 5, 11, 13, 15, 17, 19, 21}))
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			findMedianSortedArrays([]int{1, 3}, []int{2})
			findMedianSortedArrays([]int{1, 2}, []int{3, 4})
			findMedianSortedArrays([]int{1}, []int{})
			findMedianSortedArrays([]int{}, []int{1})
			findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{6})
			findMedianSortedArrays([]int{1, 3, 5, 7, 9, 11}, []int{1000})
		}
	})
}
