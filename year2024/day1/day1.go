package day1

import (
	"fmt"
	"strings"

	"github.com/Drozd0f/tools"
)

func parseInput(lines []string) ([]int64, []int64) {
	parsedLeft := make([]int64, len(lines))
	parsedRight := make([]int64, len(lines))
	for idx, line := range lines {
		parsedSlice := strings.Split(line, part1InputSeparator)
		if len(parsedSlice) != 2 {
			panic(fmt.Sprintf("parse slice: invalid line: %s", line))
		}

		parsedLeft[idx] = tools.MustInt64FromString(parsedSlice[0])
		parsedRight[idx] = tools.MustInt64FromString(parsedSlice[1])
	}

	return parsedLeft, parsedRight
}
