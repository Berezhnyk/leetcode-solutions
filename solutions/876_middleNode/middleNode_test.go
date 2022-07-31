package middlenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	r := middleNode(l1)
	for l2.Next != nil {
		assert.Equal(t, r.Val, l2.Val)
		r = r.Next
		l2 = l2.Next
	}
}

func Test_middleNode2(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}
	l2 := &ListNode{
		Val: 0,
	}

	r := middleNode(l1)
	for l2.Next != nil {
		assert.Equal(t, r.Val, l2.Val)
		r = r.Next
		l2 = l2.Next
	}
}

func Test_middleNode3(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 7,
				},
			},
		},
	}

	r := middleNode(l1)
	for l2.Next != nil {
		assert.Equal(t, r.Val, l2.Val)
		r = r.Next
		l2 = l2.Next
	}
}

func Benchmark_middleNode(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			middleNode(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			})

			middleNode(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			})

			middleNode(&ListNode{
				Val: 1,
			})
		}
	})
}
