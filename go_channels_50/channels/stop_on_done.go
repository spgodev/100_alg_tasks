package channels

func StopOnDone(done <-chan struct{}, in <-chan int) []int {
	// TODO: читать числа из in, пока in не закрыт или пока не закрыли done.
	result := make([]int, 0)
	for {
		select {
		case <-done:
			return result
		case value, ok := <-in:
			if !ok {
				return result
			}
			result = append(result, value)
		}
	}
}
