package channels

func MapChannel(in <-chan int, fn func(int) int) <-chan int {
	// TODO: применить fn к каждому значению канала.
	out := make(chan int, 100)
	go func() {
		defer close(out)
		for value := range in {
			out <- fn(value)
		}
	}()
	return out
}
