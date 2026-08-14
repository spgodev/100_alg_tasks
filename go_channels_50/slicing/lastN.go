package slicing

// LastN возвращает последние n элементов.
// Если n <= 0 — вернуть пустой слайс.
// Если n больше длины nums — вернуть весь nums.

func LastN(nums []int, n int) []int {
	// TODO
	if n <= 0 {
		return []int{}
	}
	if n >= len(nums) {
		return nums
	}
	return nums[len(nums)-n:]
}
