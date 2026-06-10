package mapslice

// HasKey сообщает, есть ли key в map.
func HasKey(m map[string]int, key string) bool {
	// TODO: реализовать функцию.
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
