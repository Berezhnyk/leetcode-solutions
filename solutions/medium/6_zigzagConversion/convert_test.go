package zigzagconversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "A", convert("A", 1))
}

func Benchmark_convert(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			convert("PAYPALISHIRING", 3)
		}
	})
}
