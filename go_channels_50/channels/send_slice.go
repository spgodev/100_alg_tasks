package channels

func SendSlice(nums []int) <-chan int {
	// TODO: отправить все числа из слайса в канал и закрыть канал.
	ch := make(chan int, len(nums))
	for _, num := range nums {
		ch <- num
	}
	close(ch)
	return ch
}
