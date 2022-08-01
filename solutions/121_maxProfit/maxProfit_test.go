package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
	assert.Equal(t, 2, maxProfit([]int{2, 4, 1}))
}

func Benchmark_maxProfit(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			maxProfit([]int{7, 1, 5, 3, 6, 4})
			maxProfit([]int{7, 6, 4, 3, 1})
		}
	})
}
