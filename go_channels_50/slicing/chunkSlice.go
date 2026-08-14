package slicing

// ChunkSlice разбивает nums на чанки размера size.
// Последний чанк может быть меньше.
// Если size <= 0 — вернуть пустой слайс.
//
// ChunkSlice([]int{1,2,3,4,5}, 2)
// = [][]int{{1,2}, {3,4}, {5}}

func ChunkSlice(nums []int, size int) [][]int {
	// TODO
	if size <= 0 {
		return [][]int{}
	}
	result := make([][]int, 0, len(nums))
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}

		chunk := nums[i:end]
		result = append(result, chunk)
	}
	return result
}
