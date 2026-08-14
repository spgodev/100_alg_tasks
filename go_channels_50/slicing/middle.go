package slicing

// Middle возвращает слайс без первого и последнего элемента.
// Если длина меньше 3 — вернуть пустой слайс.
//
// Middle([]int{1,2,3,4}) = []int{2,3}
func Middle(nums []int) []int {
	// TODO
	if len(nums) < 3 {
		return []int{}
	}
	return nums[1 : len(nums)-1]
}
