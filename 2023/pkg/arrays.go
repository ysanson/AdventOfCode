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

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func CreateSlice[T any](len int, defaultValue T) []T {
	arr := make([]T, len)
	for idx := range arr {
		arr[idx] = defaultValue
	}
	return arr
}
