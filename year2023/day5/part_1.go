package day5

import (
	"strconv"
	"strings"
)

type source struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func newSource(drs, srs, rl string) source {
	var (
		s   source
		err error
	)

	s.destinationRangeStart, err = strconv.Atoi(drs)
	if err != nil {
		panic(err)
	}

	s.sourceRangeStart, err = strconv.Atoi(srs)
	if err != nil {
		panic(err)
	}

	s.rangeLength, err = strconv.Atoi(rl)
	if err != nil {
		panic(err)
	}

	return s
}

func CalculatePart1(lines []string) int {
	rawSeeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")

	return calculateMinLocation(parceSeeds(rawSeeds), parseSourcesMap(lines[2:]))
}

func parceSeeds(rawSeeds []string) []int {
	var err error

	seeds := make([]int, len(rawSeeds))
	for idx, rawSeed := range rawSeeds {
		seeds[idx], err = strconv.Atoi(rawSeed)
		if err != nil {
			panic(err)
		}
	}

	return seeds
}

func parseSourcesMap(lines []string) [][]source {
	sourcesMap := make([][]source, 0, 7)
	sources := make([]source, 0, 40)

	for _, line := range lines {
		rawSource := strings.Split(line, " ")
		if len(rawSource) != 3 {
			if line == "" {
				sourcesMap = append(sourcesMap, sources)
				sources = make([]source, 0, 40)
			}

			continue
		}

		sources = append(sources, newSource(rawSource[0], rawSource[1], rawSource[2]))
	}

	return append(sourcesMap, sources)
}

func calculateMinLocation(seeds []int, sourcesMap [][]source) int {
	minLocation := seeds[0]
	for _, seed := range seeds {
		if location := calculateLocation(seed, sourcesMap); location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func calculateLocation(location int, sourcesMap [][]source) int {
	for _, sources := range sourcesMap {
		for _, s := range sources {
			if s.sourceRangeStart <= location && location <= s.sourceRangeStart+s.rangeLength {
				location = location - s.sourceRangeStart + s.destinationRangeStart

				break
			}
		}
	}

	return location
}
