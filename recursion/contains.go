package recursion

// Contains рекурсивно проверяет, есть ли target в слайсе.
// Contains([]int{1,2,3}, 2) = true
// Contains([]int{1,2,3}, 5) = false
func Contains(nums []int, target int) bool {
	// TODO
	if len(nums) == 0 {
		return false
	}
	if nums[0] == target {
		return true
	}
	return Contains(nums[1:], target)
}
