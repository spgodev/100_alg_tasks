package channels

func TakeN(in <-chan int, n int) <-chan int {
	// TODO: взять первые n значений из канала.
	out := make(chan int, 100)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			value, ok := <-in
			if !ok {
				return
			}
			out <- value
		}
	}()

	return out
}
