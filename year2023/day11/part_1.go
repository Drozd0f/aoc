package day11

import (
	"fmt"
)

func parceImage(lines []string) [][]string {
	image := make([][]string, len(lines))
	for lineIdx, line := range lines {
		imageLine := make([]string, len(line))
		for charIdx, char := range line {
			imageLine[charIdx] = string(char)
		}

		image[lineIdx] = imageLine
	}

	return image
}

func CalculatePartOne(lines []string) {
	image := parceImage(lines)
	galaxies := newGalaxies(image)
	lineExpansion, colExpansion := newExpansions(image, 2)
	fmt.Println(calculatePath(galaxiesWithExpansions(galaxies, lineExpansion, colExpansion)))
}

func calculatePath(galaxies []galaxy) int64 {
	var path int64
	for idx, galaxyFrom := range galaxies {
		for _, galaxyTo := range galaxies[idx+1:] {
			path += galaxyFrom.Path(galaxyTo)
		}
	}

	return path
}
