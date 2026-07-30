package channels

import "context"

func MapWithContext(ctx context.Context, in <-chan int, fn func(int) int) <-chan int {
	// TODO: map-pipeline, который прекращает работу при отмене context.
	return nil
}
