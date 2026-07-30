package channels

func RepeatValue(value int, n int) <-chan int {
	// TODO: отправить value n раз в канал и закрыть канал.
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			ch <- value
		}
	}()
	return ch
}
