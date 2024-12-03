package day3

import (
	"regexp"
	"strings"
)

var invalidMulRe = regexp.MustCompile(`don't\(.*?(do\(|$)`)

func SolutionForPartTwo(lines []string) int64 {
	return calculateMul(invalidMulRe.ReplaceAllString(strings.Join(lines, ""), ""))
}
