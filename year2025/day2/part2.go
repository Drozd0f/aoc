package day2

import (
	"strconv"
	"strings"

	"github.com/Drozd0f/tools"
)

func Part2String(seqs [][2]int64) int64 {
	var res int64
	for _, seq := range seqs {
		for i := seq[0]; i <= seq[1]; i++ {
			s := strconv.Itoa(int(i))
			lenS := len(s)
			for idx := range lenS / 2 {
				subS := s[:idx+1]
				repetCount := strings.Count(s, subS)
				expectedRepetCount := lenS / (idx + 1)
				subLen := len(subS) * (expectedRepetCount)
				if repetCount == expectedRepetCount && subLen == lenS {
					res += i
					break
				}
			}
		}
	}

	return res
}

func Part2Int(seqs [][2]int64) int64 {
	var res int64
	for _, seq := range seqs {
		for i := seq[0]; i <= seq[1]; i++ {
			lenI := tools.LenInt(i)
			for idx := range lenI / 2 {
				subI := tools.CutFromInt(i, 0, idx+1)
				repetCount := tools.CountFromInt(i, subI)
				expectedRepetCount := lenI / (idx + 1)
				subLen := tools.LenInt(subI) * expectedRepetCount
				if repetCount == expectedRepetCount && subLen == lenI {
					res += i
					break
				}
			}
		}
	}

	return res
}
