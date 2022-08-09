package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion1(t *testing.T) {
	isBadVersion = func(version int) bool {
		return version > 5
	}

	assert.Equal(t, 6, firstBadVersion(10))
}

func Test_firstBadVersion2(t *testing.T) {
	isBadVersion = func(version int) bool {
		return version > 100
	}

	assert.Equal(t, 101, firstBadVersion(100000))
}

func Test_firstBadVersion3(t *testing.T) {
	isBadVersion = func(version int) bool {
		return version > 0
	}

	assert.Equal(t, 1, firstBadVersion(1))
}

func Benchmark_isSubsequence(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			firstBadVersion(1)
		}
	})
}
