package mergetwolists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode = nil

	if (*list1).Val <= (*list2).Val {
		head = &ListNode{
			Val: (*list1).Val,
		}
		list1 = (*list1).Next
	} else {
		head = &ListNode{
			Val: (*list2).Val,
		}
		list2 = (*list2).Next
	}

	current := head
	for {
		if list1 == nil {
			(*current).Next = list2
			return head
		}

		if list2 == nil {
			(*current).Next = list1
			return head
		}

		if (*list1).Val <= (*list2).Val {
			(*current).Next = &ListNode{
				Val: (*list1).Val,
			}
			list1 = (*list1).Next

		} else {
			(*current).Next = &ListNode{
				Val: (*list2).Val,
			}
			list2 = (*list2).Next
		}
		current = (*current).Next
	}
}
