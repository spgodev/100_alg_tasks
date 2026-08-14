package slicing

// InsertAt возвращает новый слайс,
// где value вставлен по индексу index.
// Если index < 0 — вставить в начало.
// Если index > len(nums) — вставить в конец.
//
// InsertAt([]int{1,2,4}, 2, 3) = []int{1,2,3,4}

func InsertAt(nums []int, index int, value int) []int {
	// TODO
	result := make([]int, 0, len(nums)+1)
	if index < 0 {
		index = 0
	}
	if index > len(nums) {
		index = len(nums)
	}
	result = append(result, nums[:index]...)
	result = append(result, value)
	result = append(result, nums[index:]...)
	return result
}
