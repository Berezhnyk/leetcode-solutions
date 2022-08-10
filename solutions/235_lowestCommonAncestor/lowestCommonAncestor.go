package lowestcommonancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal := p.Val
	qVal := q.Val
	node := root
	for {
		currentVal := node.Val
		if pVal > currentVal && qVal > currentVal {
			node = node.Right
		} else if pVal < currentVal && qVal < currentVal {
			node = node.Left
		} else {
			return node
		}
	}
}
