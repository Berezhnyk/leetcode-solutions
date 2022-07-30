package mergetwolists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists1(t *testing.T) {
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
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	r := mergeTwoLists(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}

func Test_mergeTwoLists2(t *testing.T) {
	var l1 *ListNode = nil
	var l2 *ListNode = nil
	var l3 *ListNode = nil

	r := mergeTwoLists(l1, l2)
	assert.Equal(t, r, l3)
}

func Test_mergeTwoLists3(t *testing.T) {
	var l1 *ListNode = nil
	l2 := &ListNode{
		Val: 0,
	}
	l3 := &ListNode{
		Val: 0,
	}

	r := mergeTwoLists(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}
func Test_mergeTwoLists4(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}
	var l2 *ListNode = nil

	l3 := &ListNode{
		Val: 0,
	}

	r := mergeTwoLists(l1, l2)
	for l3.Next != nil {
		assert.Equal(t, r.Val, l3.Val)
		r = r.Next
		l3 = l3.Next
	}
}
