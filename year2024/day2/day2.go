package day2

import (
	"log"
	"strings"

	"github.com/Drozd0f/tools"
)

const separator = " "

const (
	lowerDif int64 = 1
	upperDif int64 = 3
)

func parseInput(lines []string) [][]int64 {
	result := make([][]int64, len(lines))
	for idx, line := range lines {
		parsedLine := strings.Split(line, separator)
		if len(parsedLine) == 0 {
			log.Fatalf("line %d is empty", idx)
		}

		result[idx] = make([]int64, len(parsedLine))
		for lineIdx, s := range parsedLine {
			result[idx][lineIdx] = tools.MustInt64FromString(s)
		}
	}

	return result
}

func between(n, left, right int64) bool {
	return left <= n && n <= right
}

func calculateLineResult(line []int64) int64 {
	var lineResult int64
	for idx := 1; idx < len(line); idx++ {
		dif := line[idx-1] - line[idx]

		// Any two adjacent levels differ by at least one and at most three.
		if lowerDif > tools.Abs(dif) || tools.Abs(dif) > upperDif {
			return 0
		}

		if (lineResult > 0 && dif < 0) || (lineResult < 0 && dif > 0) {
			return 0
		}

		lineResult += dif
	}

	return lineResult
}
