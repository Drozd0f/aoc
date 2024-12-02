package day2

import (
	"github.com/Drozd0f/tools"
)

func SolutionForPartOne(lines []string) int64 {
	parsedLines := parseInput(lines)

	var res int64
	for _, line := range parsedLines {
		lineResult := calculateLineResult(line)

		combineLength := int64(len(line)) - 1
		if between(tools.Abs(lineResult), lowerDif*combineLength, upperDif*combineLength) {
			res++
		}
	}

	return res
}
