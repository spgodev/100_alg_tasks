package mapslice

// RenameMapKey возвращает копию m, где значение oldKey перенесено в newKey. Если oldKey нет, map не меняется.
func RenameMapKey(m map[string]int, oldKey string, newKey string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)

	for k, v := range m {
		result[k] = v
	}
	value, ok := result[oldKey]
	if !ok {
		return result
	}

	delete(result, oldKey)
	result[newKey] = value
	return result
}
