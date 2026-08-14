package slicing

// DropFirst возвращает слайс без первого элемента.
// Если слайс пустой — вернуть пустой слайс.
func DropFirst(nums []int) []int {
	// TODO
	if len(nums) == 0 {
		return []int{}
	}
	return nums[1:]
}
