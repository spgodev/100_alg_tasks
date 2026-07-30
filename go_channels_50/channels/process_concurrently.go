package channels

import "sync"

func ProcessConcurrently(nums []int, fn func(int) int) []int {
	// TODO: обработать каждый элемент параллельно и вернуть результаты в исходном порядке.
	result := make([]int, len(nums))

	wg := sync.WaitGroup{}
	for i := range nums {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			result[index] = fn(nums[index])
		}(i)
	}
	wg.Wait()
	return result
}
