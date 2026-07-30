package recursion

type Node struct {
	Value    int
	Children []*Node
}

// TreeSum рекурсивно считает сумму всех Value в дереве.
// Если root == nil — вернуть 0.
func TreeSum(root *Node) int {
	// TODO
	if root == nil {
		return 0
	}
	result := root.Value
	for _, child := range root.Children {
		result += TreeSum(child)
	}
	return result
}
