package pkg

func MapContains[T, M comparable](m map[T]M, value M) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func MapCountOccurrences[T, M comparable](m map[T]M, value M) int {
	occurrence := 0
	for _, v := range m {
		if v == value {
			occurrence++
		}
	}
	return occurrence
}
