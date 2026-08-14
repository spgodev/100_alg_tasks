package slicing

// DropLast возвращает слайс без последнего элемента.
// Если слайс пустой — вернуть пустой слайс.
func DropLast(nums []int) []int {
	// TODO
	if len(nums) == 0 {
		return []int{}
	}
	return nums[:len(nums)-1]
}
