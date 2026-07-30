package recursion

// MaxInSlice рекурсивно находит максимум в слайсе.
// Если слайс пустой — вернуть 0.
func MaxInSlice(nums []int) int {
	// TODO
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxRest := MaxInSlice(nums[1:])
	if nums[0] > maxRest {
		return nums[0]
	}
	return maxRest
}
