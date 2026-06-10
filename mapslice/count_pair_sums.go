package mapslice

// CountPairSums возвращает количество пар индексов i < j, для которых src[i] + src[j] == target.
func CountPairSums(src []int, target int) int {
	// TODO: реализовать функцию.
	seen := make(map[int]int)
	result := 0

	for _, v := range src {
		need := target - v
		result += seen[need]
		seen[v]++
	}
	return result
}
