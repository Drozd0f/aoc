package day3

import "strings"

func SolutionForPartOne(lines []string) int64 {
	return calculateMul(strings.Join(lines, ""))
}
