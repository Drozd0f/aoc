package day6

import (
	"strconv"
	"strings"
)

func parceRacePart2(lines []string) race {
	rawDurations := strings.Split(lines[0], ":")
	rawDistances := strings.Split(lines[1], ":")

	duration := strings.Join(numberRegex.FindAllString(rawDurations[1], -1), "")
	distance := strings.Join(numberRegex.FindAllString(rawDistances[1], -1), "")

	intDuration, err := strconv.Atoi(duration)
	if err != nil {
		panic(err)
	}

	intDistance, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}

	return race{
		duration: int64(intDuration),
		distance: int64(intDistance),
	}
}

func CalculatePart2(lines []string) int64 {
	return calculateWinningWaysPart2(parceRacePart2(lines))
}

func calculateWinningWaysPart2(r race) int64 {
	var winningWay int64
	for t := r.duration / 2; t >= 1; t-- {
		if raceDuration := r.duration - t; r.distance < raceDuration*t {
			if raceDuration == t {
				winningWay++

				continue
			}

			winningWay += 2

			continue
		}

		break
	}

	return winningWay
}
