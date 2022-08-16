package lengthoflongestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("aabaab!bb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
	assert.Equal(t, 2, lengthOfLongestSubstring("aab"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcab"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lengthOfLongestSubstring("abcabcab")
		}
	})
}
