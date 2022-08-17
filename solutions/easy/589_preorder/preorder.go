package preorder

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}
	result := []int{root.Val}

	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}

	return result
}
