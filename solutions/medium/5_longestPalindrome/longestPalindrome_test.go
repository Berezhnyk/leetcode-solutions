package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "kek", longestPalindrome("suslolpopkek"))
	assert.Equal(t, "lolpoplol", longestPalindrome("suslolpoplolkek"))
	assert.Equal(t, "lolpoplol", longestPalindrome("suslolpoplolkektetate"))
	assert.Equal(t, "aba", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}

func Benchmark_longestPalindrome(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			longestPalindrome("suslolpopkek")
			longestPalindrome("suslolpoplolkek")
			longestPalindrome("suslolpoplolkektetate")
			longestPalindrome("babad")
			longestPalindrome("cbbd")
		}
	})
}
