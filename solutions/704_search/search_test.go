package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func Benchmark_search(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			search([]int{-1, 0, 3, 5, 9, 12}, 9)
			search([]int{-1, 0, 3, 5, 9, 12}, 2)
		}
	})
}
