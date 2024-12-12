package pkg

func MapContains[T, M comparable](m map[T]M, value M) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func SetContainsAll[T, M comparable](set map[T]M, key ...T) bool {
	for _, key := range key {
		if _, ok := set[key]; !ok {
			return false
		}

	}
	return true
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
