package mapslice

// CanMakePairSum возвращает true, если существуют два разных индекса с суммой target.
func CanMakePairSum(src []int, target int) bool {
	// TODO: реализовать функцию.
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i]+src[j] == target {
				return true
			}
		}
	}
	return false
}

/*seen := make(map[int]bool)

for _, value := range src {
need := target - value

if seen[need] {
return true
}

seen[value] = true
}

return false
}
*/
