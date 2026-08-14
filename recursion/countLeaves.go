package recursion

// CountLeaves возвращает количество листьев в дереве.
// Если root == nil — вернуть 0.
//
// Пример:
//      1
//    / | \
//   2  3  4
//
// Листья: 2, 3, 4
// CountLeaves(root) = 3

func CountLeaves(root *TreeNode) int {
	// TODO
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	counter := 0
	for _, child := range root.Children {
		counter += CountLeaves(child)
	}
	return counter
}
