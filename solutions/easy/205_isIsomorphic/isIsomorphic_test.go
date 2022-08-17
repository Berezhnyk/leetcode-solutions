package isisomorphic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	assert.Equal(t, false, isIsomorphic("babc", "baba"))
	assert.Equal(t, true, isIsomorphic("egg", "add"))
	assert.Equal(t, false, isIsomorphic("foo", "bar"))
	assert.Equal(t, true, isIsomorphic("paper", "title"))
}

func Benchmark_isIsomorphic(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isIsomorphic("egg", "add")
			isIsomorphic("foo", "bar")
			isIsomorphic("paper", "title")
		}
	})
}
