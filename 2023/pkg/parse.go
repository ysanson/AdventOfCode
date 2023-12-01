package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}

func ParseIntList(s, sep string) []int {
	lines := strings.Split(s, sep)
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		Check(err)
		list[i] = nb
	}
	return list
}

func ParseIntMap(s, sep string) map[int]int {
	m := make(map[int]int)
	for i, line := range strings.Split(s, sep) {
		nb, err := strconv.Atoi(line)
		Check(err)
		m[i] = nb
	}
	return m
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func MustScanf(line, format string, a ...interface{}) {
	n, err := fmt.Sscanf(line, format, a...) // don't forget to set parseCount
	Check(err)
	if n != len(a) {
		panic(fmt.Errorf("%d args expected in scanf, got %d", len(a), n))
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
