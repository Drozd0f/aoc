package day5

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

func CalculatePart2(lines []string) int {
	rawSeeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")

	return calculateMinLocationPart2(parseSeedsRange(rawSeeds), parseSourcesMap(lines[2:]))
}

func parseSeedsRange(rawSeeds []string) [][2]int {
	seeds := parceSeeds(rawSeeds)
	if len(seeds)%2 != 0 {
		panic(fmt.Sprintf("invalid seeds lenth: %d", len(seeds)))
	}

	seedsRange := make([][2]int, 0, len(seeds)/2)
	for idx := 0; idx < len(seeds); idx += 2 {
		seedsRange = append(seedsRange, [2]int{seeds[idx], seeds[idx+1]})
	}

	return seedsRange
}

func calculateMinLocationPart2(seedsRange [][2]int, sourcesMap [][]source) int {
	var wg sync.WaitGroup
	minLocations := asyncMinLocations{v: make([]int, 0, len(seedsRange))}

	for _, sr := range seedsRange {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			minLocation := start
			for seed := start; seed < start+end; seed++ {
				if location := calculateLocation(seed, sourcesMap); location < minLocation {
					minLocation = location
				}
			}

			minLocations.Append(minLocation)
		}(sr[0], sr[1])
	}

	wg.Wait()

	return slices.Min(minLocations.v)
}

type asyncMinLocations struct {
	mu sync.Mutex
	v  []int
}

func (ml *asyncMinLocations) Append(v int) {
	ml.mu.Lock()

	ml.v = append(ml.v, v)

	ml.mu.Unlock()
}
