package mapslice

// MostFrequentInt возвращает самое частое число. При равенстве частот возвращает меньшее число. Для пустого слайса возвращает 0.
func MostFrequentInt(src []int) int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return 0
	}

	maxFreq := 0
	bestValue := src[0]
	freqCounter := make(map[int]int)
	for _, value := range src {
		freqCounter[value]++
	}
	for value, freq := range freqCounter {
		if freq > maxFreq || freq == maxFreq && value < bestValue {
			maxFreq = freq
			bestValue = value
		}
	}

	return bestValue
}
