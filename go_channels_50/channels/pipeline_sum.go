package channels

func PipelineSum(in <-chan int) <-chan int {
	// TODO: отправить сумму всех входных чисел одним значением и закрыть канал.
	out := make(chan int, 100)
	sum := 0
	go func() {
		defer close(out)
		for value := range in {
			sum += value
		}
		out <- sum
	}()

	return out
}
