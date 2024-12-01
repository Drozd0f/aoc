package day9

import (
	"fmt"
	"strings"

	"github.com/Drozd0f/tools"
)

func parceSequence(lines []string) [][]int64 {
	sequence := make([][]int64, len(lines))
	for lineIdx, line := range lines {
		rawSequence := strings.Split(line, " ")
		seq := make([]int64, len(rawSequence))
		for seqIdx, rawSeq := range rawSequence {
			seq[seqIdx] = tools.MustInt64FromString(rawSeq)
		}

		sequence[lineIdx] = seq
	}

	return sequence
}

func CalculatePartOne(lines []string) {
	fmt.Println(calculatePartOne(parceSequence(lines)))
}

func calculatePartOne(sequence [][]int64) int64 {
	var historyValues int64
	for _, history := range sequence {
		historyValues += calculateHistoryValue(history)
	}

	return historyValues
}

func calculateHistoryValue(sequence []int64) int64 {
	var isAllZero = true

	lasIdx := len(sequence) - 1
	newSequence := make([]int64, lasIdx)
	for idx := 0; idx < lasIdx; idx++ {
		newSeq := sequence[idx+1] - sequence[idx]
		if newSeq != 0 && isAllZero {
			isAllZero = false
		}

		newSequence[idx] = newSeq
	}

	if isAllZero {
		return sequence[lasIdx]
	}

	return sequence[lasIdx] + calculateHistoryValue(newSequence)
}
