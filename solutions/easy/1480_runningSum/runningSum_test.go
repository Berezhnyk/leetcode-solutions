package runningsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 10}, runningSum([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, runningSum([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, []int{3, 4, 6, 16, 17}, runningSum([]int{3, 1, 2, 10, 1}))
}

func Benchmark_runningSum(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runningSum([]int{1, 2, 3, 4})
			runningSum([]int{1, 1, 1, 1, 1})
			runningSum([]int{3, 1, 2, 10, 1})
		}
	})
}
