package day11

import (
	"slices"

	"github.com/Drozd0f/tools"
)

type expansion struct {
	location int64
	size     int64
}

func newExpansions(image [][]string, size int64) ([]expansion, []expansion) {
	lineLen := len(image[0])
	colLen := len(image)
	lineExpansions := make([]expansion, 0, lineLen)
	for lineIdx := 0; lineIdx < colLen; lineIdx++ { // expansion lines
		if lineLen == tools.Count(image[lineIdx], emptySpace) {
			lineExpansions = append(lineExpansions, expansion{
				location: int64(lineIdx),
				size:     size - 1,
			})
		}
	}

	columnExpansions := make([]expansion, 0, colLen)
	for colIdx := 0; colIdx < lineLen; colIdx++ {
		column := make([]string, colLen)
		for lineIdx := 0; lineIdx < colLen; lineIdx++ {
			column[lineIdx] = image[lineIdx][colIdx]
		}

		if colLen == tools.Count(column, emptySpace) {
			columnExpansions = append(columnExpansions, expansion{
				location: int64(colIdx),
				size:     size - 1,
			})
		}
	}

	return slices.Clip(lineExpansions), slices.Clip(columnExpansions)
}
