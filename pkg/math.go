package pkg

import (
	"math"
	"math/bits"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type nullableInt struct {
	value int
}

func Max(values ...int) int {
	if len(values) == 0 {
		panic("no value in max function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || max.value < value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func Min(values ...int) int {
	if len(values) == 0 {
		panic("no value in min function")
	}

	var min *nullableInt
	for _, value := range values {
		if min == nil || min.value > value {
			min = &nullableInt{value}
		}
	}
	return min.value
}

func Sum(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func Multiply(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 1
	for _, value := range values {
		sum *= value
	}
	return sum
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

func LcmAll(a int, bs ...int) int {
	result := a
	for _, b := range bs {
		result = LCM(result, b)
	}

	return result
}

func TransposeMatrix[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))
	if n > len(set) {
		n = len(set)
	}
	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}
		var subset []T
		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func IsInteger(num float64) bool {
	return math.Mod(num, 1) == 0
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func UintPow(n, m uint64) uint64 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; uint64(i) <= m; i++ {
		result *= n
	}
	return result
}
