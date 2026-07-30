package channels

func GracefulShutdownWorkers(jobs <-chan int, workers int, fn func(int) int) <-chan int {
	// TODO: worker pool, который закрывает results после завершения всех workers.
	return nil
}
