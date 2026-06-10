package mapslice

// EqualStringIntMaps сообщает, равны ли две map по набору ключей и значений.
func EqualStringIntMaps(a map[string]int, b map[string]int) bool {
	// TODO: реализовать функцию.
	//return reflect.DeepEqual(a, b)
	if len(a) != len(b) {
		return false
	}

	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok {
			return false
		}

		if valueA != valueB {
			return false
		}
	}

	return true
}
