package execute

import (
	"fmt"
	"os"
	"testing"
)

type TestCases []TestCase

// TestCase holds an input and its answers. An expected value of "" is not
// checked, and a case with an empty Input is skipped.
type TestCase struct {
	Input                        string
	ExpectedPart1, ExpectedPart2 interface{}
}

// mismatch describes a wrong answer, or returns "" if the answer is accepted.
func mismatch(part int, got, expected interface{}) string {
	gotStr, expectedStr := fmt.Sprint(got), fmt.Sprint(expected)
	if expectedStr == "" || gotStr == expectedStr {
		return ""
	}
	return fmt.Sprintf(" - PART%d: Got %s but expected %s", part, gotStr, expectedStr)
}

func (t TestCases) Run(fn func(string) (interface{}, interface{}), hideInput bool) {
	for _, test := range t {
		if test.Input == "" {
			continue
		}
		part1, part2 := fn(test.Input)
		failure1, failure2 := mismatch(1, part1, test.ExpectedPart1), mismatch(2, part2, test.ExpectedPart2)

		if (failure1 != "" || failure2 != "") && !hideInput {
			fmt.Println("Input\n", test.Input)
		}
		for _, failure := range []string{failure1, failure2} {
			if failure != "" {
				fmt.Println(failure)
				os.Exit(1)
			}
		}
	}
}

// RunT runs each case as a subtest, so one failure does not hide the others.
func (t TestCases) RunT(tt *testing.T, fn func(string) (interface{}, interface{})) {
	for i, test := range t {
		tt.Run(fmt.Sprintf("case%d", i+1), func(st *testing.T) {
			if test.Input == "" {
				st.Skip("empty input")
			}
			part1, part2 := fn(test.Input)
			if failure := mismatch(1, part1, test.ExpectedPart1); failure != "" {
				st.Error(failure)
			}
			if failure := mismatch(2, part2, test.ExpectedPart2); failure != "" {
				st.Error(failure)
			}
		})
	}
}
