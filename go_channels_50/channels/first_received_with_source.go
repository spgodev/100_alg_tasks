package channels

func FirstReceivedWithSource(a <-chan int, b <-chan int) (int, string) {
	// TODO: вернуть значение и имя канала: "a" или "b".
	select {
	case value := <-a:
		return value, "a"
	case value := <-b:
		return value, "b"
	}
}
