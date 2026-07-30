package recursion

// SumSlice рекурсивно считает сумму элементов слайса.
// SumSlice([]int{1,2,3}) = 6
// Для пустого слайса вернуть 0.
func SumSlice(nums []int) int {
	// TODO
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumSlice(nums[1:])
}
