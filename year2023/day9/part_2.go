package day9

import (
	"fmt"
	"slices"
)

func CalculatePartTwo(lines []string) {
	fmt.Println(calculatePartTwo(parceSequence(lines)))
}

func calculatePartTwo(sequence [][]int64) int64 {
	var historyValues int64
	for _, history := range sequence {
		slices.Reverse(history)
		res := calculateHistoryValue(history)

		historyValues += res
	}

	return historyValues
}
