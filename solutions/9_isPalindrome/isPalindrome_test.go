package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(1234321))
	assert.Equal(t, true, isPalindrome(2468642))
	assert.Equal(t, true, isPalindrome(7))
	assert.Equal(t, false, isPalindrome(-1234321))
	assert.Equal(t, false, isPalindrome(123456543210))
}

func Benchmark_isPalindrome(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			isPalindrome(1)
			isPalindrome(-1234321)
			isPalindrome(1234321)
			isPalindrome(-1234321)
			isPalindrome(424)
			isPalindrome(123456543210)
		}
	})
}
