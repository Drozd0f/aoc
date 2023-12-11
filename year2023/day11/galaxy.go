package day11

import "math"

const (
	emptySpace     = "."
	galaxiesSymbol = "#"
)

type galaxy struct {
	location [2]int64
}

func newGalaxies(image [][]string) []galaxy {
	var galaxies []galaxy
	for idx, line := range image {
		for jdx, object := range line {
			if galaxiesSymbol == object {
				galaxies = append(galaxies, galaxy{
					location: [2]int64{int64(idx), int64(jdx)},
				})
			}
		}
	}

	return galaxies
}

func (g galaxy) Path(other galaxy) int64 {
	return int64(
		math.Abs(float64(g.location[0]-other.location[0])) +
			math.Abs(float64(g.location[1]-other.location[1])),
	)
}

func galaxiesWithExpansions(galaxies []galaxy, lineExpansions, columnExpansions []expansion) []galaxy {
	lineGalaxies := make([]galaxy, len(galaxies))
	for idx, g := range galaxies {
		var expansionSize int64
		for _, lineExpansion := range lineExpansions {
			if g.location[0] < lineExpansion.location {
				break
			}

			expansionSize += lineExpansion.size
		}

		g.location[0] += expansionSize
		lineGalaxies[idx] = g
	}

	colGalaxies := make([]galaxy, len(galaxies))
	for idx, g := range lineGalaxies {
		var expansionSize int64
		for _, columnExpansion := range columnExpansions {
			if g.location[1] < columnExpansion.location {
				break
			}

			expansionSize += columnExpansion.size
		}

		g.location[1] += expansionSize
		colGalaxies[idx] = g
	}

	return colGalaxies
}
