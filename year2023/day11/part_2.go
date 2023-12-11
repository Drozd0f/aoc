package day11

import (
	"fmt"
)

func CalculatePartTwo(lines []string) {
	// 10077850
	image := parceImage(lines)
	galaxies := newGalaxies(image)
	lineExpansion, colExpansion := newExpansions(image, 1000000)
	fmt.Println(calculatePath(galaxiesWithExpansions(galaxies, lineExpansion, colExpansion)))
}
