package recursion

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

// CountNodes возвращает количество всех узлов в дереве.
// Если root == nil — вернуть 0.
//
// Пример:
//
//	   1
//	 / | \
//	2  3  4
//
// CountNodes(root) = 4
func CountNodes(root *TreeNode) int {
	// TODO
	counter := 1
	if root == nil {
		return 0
	}
	for _, child := range root.Children {
		counter += CountNodes(child)
	}
	return counter
}
