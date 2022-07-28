package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	res := *addTwoNumbers(&l1, &l2)

	fmt.Println("RESULT")
	for res.Next != nil {
		fmt.Println(res.Val)
		res = *res.Next
	}
	fmt.Println(res.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	add := 0
	current := &result

	for {
		if l1 != nil || l2 != nil {
			sum := add
			if l1 != nil {
				sum += l1.Val
			}

			if l2 != nil {
				sum += l2.Val
			}
			current.Val = sum % 10
			add = sum / 10
		} else {
			if add > 0 {
				current.Next = &ListNode{Val: add}
			}
			break
		}

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

		if l1 != nil || l2 != nil {
			current.Next = &ListNode{}
			current = current.Next
		}
	}

	return &result
}
