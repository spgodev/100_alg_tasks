package mapslice

// BuildStringSet строит set из words в виде map[string]bool.
func BuildStringSet(words []string) map[string]bool {
	// TODO: реализовать функцию.
	result := make(map[string]bool)
	for _, value := range words {
		result[value] = true
	}
	return result
}
