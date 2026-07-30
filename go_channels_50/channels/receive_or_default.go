package channels

func ReceiveOrDefault(ch <-chan int, defaultValue int) int {
	// TODO: прочитать значение, если оно уже есть, иначе вернуть defaultValue.

	select {
	case val := <-ch:
		return val
	default:
		return defaultValue
	}
}
