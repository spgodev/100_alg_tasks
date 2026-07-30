package channels

func SplitEvenOdd(in <-chan int) (<-chan int, <-chan int) {
	// TODO: четные отправлять в even, нечетные в odd.
	resultEven := make(chan int, 100)
	resultOdd := make(chan int, 100)
	go func() {
		defer close(resultOdd)
		defer close(resultEven)
		for value := range in {
			if value%2 == 0 {
				resultEven <- value
			} else {
				resultOdd <- value
			}
		}
	}()
	return resultEven, resultOdd
}
