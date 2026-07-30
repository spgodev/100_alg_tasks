package channels

import "sync"

func WorkerPoolSquare(in <-chan int, workers int) <-chan int {
	// TODO: запустить workers горутин, которые возвращают квадраты чисел.
	outs := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for value := range in {
				outs <- value * value
			}
		}()

	}
	go func() {
		wg.Wait()
		close(outs)
	}()
	return outs
}
