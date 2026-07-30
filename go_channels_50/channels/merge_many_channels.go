package channels

import "sync"

func MergeManyChannels(channels []<-chan int) <-chan int {
	// TODO: объединить много каналов в один.
	out := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(input <-chan int) {
			defer wg.Done()
			for values := range input {
				out <- values
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
