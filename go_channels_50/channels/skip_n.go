package channels

func SkipN(in <-chan int, n int) <-chan int {
	// TODO: пропустить первые n значений, остальные передать дальше.
	out := make(chan int, 100)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			_, ok := <-in
			if !ok {
				return
			}
		}
		for value := range in {
			out <- value
		}
	}()

	return out
}
