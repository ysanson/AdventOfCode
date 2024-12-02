package pkg

import "iter"

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

func IsSameElement[T comparable](slice []T) bool {
	if len(slice) == 0 {
		return true
	} else {
		firstElt := slice[0]
		for _, elt := range slice {
			if elt != firstElt {
				return false
			}
		}
		return true
	}
}

func Zip[T, U any](t []T, u []U) iter.Seq2[T, U] {
    return func(yield func(T, U) bool) {
        for i := range min(len(t), len(u)) {
            if !yield(t[i], u[i]) {
                return
            }
        }
    }
}
