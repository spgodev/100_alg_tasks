package mapslice

// OnlyKeys возвращает новую map только с указанными ключами, если они есть в m.
func OnlyKeys(m map[string]int, keys []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	seen := make(map[string]bool)
	for _, key := range keys {
		seen[key] = true
	}
	for k, v := range m {
		if seen[k] {
			result[k] = v
		}
	}

	return result
}
