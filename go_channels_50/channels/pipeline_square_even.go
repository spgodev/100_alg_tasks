package channels

func PipelineSquareEven(in <-chan int) <-chan int {
	// TODO: оставить только четные числа и возвести их в квадрат.
	out := make(chan int, 100)

	go func() {
		defer close(out)
		for value := range in {
			if value%2 == 0 {
				out <- value * value
			}
		}
	}()
	return out
}
