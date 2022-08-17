package pivotindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotIndex(t *testing.T) {
	assert.Equal(t, 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, pivotIndex([]int{1, 2, 3}))
	assert.Equal(t, 0, pivotIndex([]int{2, 1, -1}))
}

func Benchmark_pivotIndex(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pivotIndex([]int{1, 7, 3, 6, 5, 6})
			pivotIndex([]int{1, 2, 3})
			pivotIndex([]int{2, 1, -1})
		}
	})
}
