package day2

import (
	"github.com/Drozd0f/tools"
)

func SolutionForPartTwo(lines []string) int64 {
	parsedLines := parseInput(lines)

	var res int64
	for _, line := range parsedLines {
		var (
			lineResult int64
			isResolve  bool
		)
		if lineResult = calculateLineResult(line); lineResult == 0 {
			newLine := make([]int64, len(line))
			for idx := 0; idx < len(line); idx++ {
				copy(newLine, line)
				if lineResult = calculateLineResult(append(newLine[:idx], newLine[idx+1:]...)); lineResult != 0 {
					if idx != 0 {
						isResolve = true
					}

					break
				}
			}
		}

		combineLength := int64(len(line) - 1)
		if isResolve {
			combineLength--
		}
		if between(tools.Abs(lineResult), lowerDif*combineLength, upperDif*combineLength) {
			res++
		}
	}

	return res
}
