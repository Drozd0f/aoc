package day1

import (
	"github.com/Drozd0f/tools"
)

const Left = "L"

const (
	minTact int64 = 0
	maxTact int64 = 100
)

func Parse(lines []string) []int64 {
	var tact int64

	tacts := make([]int64, len(lines))
	for idx, line := range lines {
		tact = tools.MustInt64FromString(line[1:])
		if string(line[0]) == Left {
			tact *= -1
		}

		tacts[idx] = tact
	}

	return tacts
}

func tactRotation(pos, tact int64) (int64, int64) {
	posTact := pos + tact
	if tools.Between(minTact-1, maxTact, posTact) {
		return 0, posTact
	}

	var rot int64
	if maxTact < tools.Abs(tact) {
		rot += tools.Abs(tact / maxTact)
		posTact = pos + tact%maxTact
	}

	if !tools.Between(minTact, maxTact+1, posTact) && pos != 0 && posTact != 0 {
		rot++
	}

	return rot, (maxTact + posTact) % maxTact
}
