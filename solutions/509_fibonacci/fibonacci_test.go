package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	assert.Equal(t, fib(0), 0)
	assert.Equal(t, fib(1), 1)
	assert.Equal(t, fib(4), 3)
	assert.Equal(t, fib(100500), 1021279670490171120)
	assert.Equal(t, fib(999999), 7006191581884273890)
}

func Benchmark_fib(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fib(0)
			fib(1)
			fib(2)
			fib(100500)
			fib(999999)
		}
	})
}
