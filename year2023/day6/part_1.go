package day6

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRegex = regexp.MustCompile(`\d+`)
)

type race struct {
	duration int64
	distance int64
}

func parceRaces(lines []string) []race {
	rawDurations := strings.Split(lines[0], ":")
	rawDistances := strings.Split(lines[1], ":")

	durations := numberRegex.FindAllString(rawDurations[1], -1)
	distances := numberRegex.FindAllString(rawDistances[1], -1)

	races := make([]race, len(durations))
	for idx, duration := range durations {
		intDuration, err := strconv.Atoi(duration)
		if err != nil {
			panic(err)
		}

		intDistance, err := strconv.Atoi(distances[idx])
		if err != nil {
			panic(err)
		}

		races[idx] = race{
			duration: int64(intDuration),
			distance: int64(intDistance),
		}
	}

	return races
}

func CalculatePart1(lines []string) int64 {
	return calculateWinningWays(parceRaces(lines))
}

func calculateWinningWays(races []race) int64 {
	winningWays := make([]int64, len(races))
	for idx, r := range races {
		for t := int64(1); t <= r.duration/2; t++ {
			if leftMult := r.duration - t; r.distance < leftMult*t {
				if leftMult == t {
					winningWays[idx]++

					continue
				}

				winningWays[idx] += 2
			}
		}
	}

	multWinningWays := winningWays[0]
	for _, winningWay := range winningWays[1:] {
		multWinningWays *= winningWay
	}

	return multWinningWays
}
