package channels

func FilterChannel(in <-chan int, predicate func(int) bool) <-chan int {
	// TODO: пропускать только значения, для которых predicate(value) == true.
	out := make(chan int, 100)
	go func() {
		defer close(out)
		for value := range in {
			if predicate(value) == true {
				out <- value
			}
		}
	}()
	return out
}
