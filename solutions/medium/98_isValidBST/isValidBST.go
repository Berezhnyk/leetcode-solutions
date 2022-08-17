package isvalidbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	return isValidNode(root, MinInt, MaxInt)
}

func isValidNode(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min {
		return false
	}

	if node.Val >= max {
		return false
	}

	return isValidNode(node.Right, node.Val, max) && isValidNode(node.Left, min, node.Val)
}
