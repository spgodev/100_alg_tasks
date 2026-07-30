package channels

import "context"

func ContextDone(ctx context.Context, ch <-chan int) (int, bool) {
	// TODO: ждать либо значение из ch, либо отмену ctx.
	return 0, false
}
