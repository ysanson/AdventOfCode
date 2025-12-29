package pkg

import (
	"iter"
	"sort"
)

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

func IntersectSorted[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool {
			return b[i] == v
		})
		if idx < len(b) && b[idx] == v {
			set = append(set, v)
		}
	}

	return set
}

func IntersectHash[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func AdjacentPairs[T any](items []T) [][2]T {
	if len(items) < 2 {
		return nil
	}

	pairs := make([][2]T, 0, len(items)-1)
	for i := 0; i < len(items)-1; i++ {
		pairs = append(pairs, [2]T{items[i], items[i+1]})
	}
	return pairs
}
