package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "pat", longestCommonPrefix([]string{"patriot", "patreon", "patrol", "pattern"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"patriot", "patreon", "dog", "patrol", "pattern"}))
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			longestCommonPrefix([]string{"flower", "flow", "flight"})
			longestCommonPrefix([]string{"dog", "racecar", "car"})
			longestCommonPrefix([]string{"patriot", "patreon", "patrol", "pattern"})
		}
	})
}
