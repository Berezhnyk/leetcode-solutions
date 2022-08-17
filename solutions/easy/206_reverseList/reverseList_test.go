package reverselist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
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

	r := reverseList(l1)
	for l2.Next != nil {
		assert.Equal(t, r.Val, l2.Val)
		r = r.Next
		l2 = l2.Next
	}
}

func Test_reverseList2(t *testing.T) {
	var l1 *ListNode = nil
	var l2 *ListNode = nil

	r := reverseList(l1)
	assert.Equal(t, r, l2)
}

func Test_reverseList3(t *testing.T) {
	var l1 *ListNode = nil
	l2 := &ListNode{
		Val: 0,
	}

	r := reverseList(l1)
	for l2.Next != nil {
		assert.Equal(t, r.Val, l2.Val)
		r = r.Next
		l2 = l2.Next
	}
}

func Benchmark_reverseList(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			reverseList(&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			})

			reverseList(&ListNode{
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

			reverseList(&ListNode{
				Val: 1,
			})

			reverseList(nil)
		}
	})
}
