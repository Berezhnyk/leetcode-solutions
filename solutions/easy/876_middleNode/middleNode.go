package middlenode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodeLength := 1
	current := head
	for current.Next != nil {
		nodeLength++
		current = current.Next
	}

	current = head
	middle := nodeLength / 2

	for i := 0; i < middle; i++ {
		current = current.Next
	}

	return current
}
