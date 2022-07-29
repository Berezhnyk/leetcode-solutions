package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	assert.Equal(t, false, isSubsequence("aaaaaa", "bbaaaa"))
	assert.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgdc"))
}

func Benchmark_isSubsequence(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isSubsequence("aaaaaa", "bbaaaa")
			isSubsequence("abc", "ahbgdc")
			isSubsequence("axc", "ahbgdc")
		}
	})
}
