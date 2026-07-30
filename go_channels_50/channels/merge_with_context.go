package channels

import "context"

func MergeWithContext(ctx context.Context, channels []<-chan int) <-chan int {
	// TODO: объединить каналы, но остановиться при отмене context.
	return nil
}
