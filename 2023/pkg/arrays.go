package pkg

func SanitizeIndex(index int, maxIndex int) int {
	if index < 0 {
		return 0
	} else if index >= maxIndex {
		return maxIndex - 1
	} else {
		return index
	}
}
