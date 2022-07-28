package nummatchingsubseq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numMatchingSubseq(t *testing.T) {
	assert.Equal(t, 3, numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	assert.Equal(t, 2, numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func Benchmark_numMatchingSubseq(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})
			numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})
		}
	})
}
