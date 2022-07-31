package minimumoperations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations(t *testing.T) {
	assert.Equal(t, 0, minimumOperations([]int{}))
	assert.Equal(t, 1, minimumOperations([]int{3}))
	assert.Equal(t, 3, minimumOperations([]int{3, 2, 4}))
	assert.Equal(t, 2, minimumOperations([]int{3, 3, 4}))
}

func Benchmark_minimumOperations(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			minimumOperations([]int{3, 3, 4})
			minimumOperations([]int{1, 2, 3})
		}
	})
}
