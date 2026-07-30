package channels

func GenerateRange(start int, end int) <-chan int {
	// TODO: сгенерировать числа от start до end включительно в канал.
	ch := make(chan int)
	go func() {
		defer close(ch)

		for i := start; i <= end; i++ {
			ch <- i
		}
	}()

	return ch
}
