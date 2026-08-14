package slicing

// RemoveAt возвращает новый слайс без элемента с индексом index.
// Если index некорректный — вернуть копию исходного слайса.
//
// RemoveAt([]int{10,20,30}, 1) = []int{10,30}

func RemoveAt(nums []int, index int) []int {
	// TODO
	result := make([]int, 0, len(nums))
	if index < 0 || index >= len(nums) {
		return append([]int{}, nums...)
	}
	result = append(result, nums[:index]...)
	result = append(result, nums[index+1:len(nums)]...)
	return result
}
