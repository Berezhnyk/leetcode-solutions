package detectcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[int][]*ListNode)
	current := head
	for current != nil && current.Next != nil {
		if nodes, ok := m[current.Val]; ok {
			if contains(nodes, current) {
				return current
			}
			m[current.Val] = append(nodes, current)
		} else {
			m[current.Val] = []*ListNode{current}
		}
		current = current.Next
	}

	return nil
}

func contains(a []*ListNode, ln *ListNode) bool {
	for _, n := range a {
		if ln == n {
			return true
		}
	}
	return false
}
