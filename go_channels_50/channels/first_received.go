package channels

func FirstReceived(a <-chan int, b <-chan int) int {
	// TODO: вернуть значение, которое пришло первым.

	select {
	case value := <-a:
		return value
	case value := <-b:
		return value
	}
}
