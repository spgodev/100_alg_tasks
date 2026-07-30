package channels

import "context"

func WorkerPoolWithContext(ctx context.Context, in <-chan int, workers int, fn func(int) int) <-chan int {
	// TODO: worker pool, который прекращает работу при отмене context.
	return nil
}
