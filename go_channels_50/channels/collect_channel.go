package channels

func CollectChannel(ch <-chan int) []int {
	// TODO: собрать все значения из канала в слайс.
	result := make([]int, 0)
	for v := range ch {
		result = append(result, v)
	}
	return result
}
