package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := handleChild([][]int{}, root, 0)

	return result
}

func handleChild(res [][]int, node *TreeNode, level int) [][]int {
	if len(res) > level {
		res[level] = append(res[level], node.Val)
	} else {
		res = append(res, []int{node.Val})
	}

	if node.Left != nil {
		res = handleChild(res, node.Left, level+1)
	}

	if node.Right != nil {
		res = handleChild(res, node.Right, level+1)
	}

	return res
}
