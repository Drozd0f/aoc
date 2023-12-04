package day4

import (
	"regexp"
	"slices"
	"strings"
)

func CalculatePart1(lines []string) int64 {
	var result int64
	for _, line := range lines {
		result += calculatePoints(extractNumbers(line))
	}

	return result
}

var extractNumbersRegex = regexp.MustCompile(`\d+`)

func extractNumbers(line string) ([]string, []string) {
	splitLine := strings.Split(strings.Split(line, ": ")[1], " | ")

	return extractNumbersRegex.FindAllString(splitLine[0], -1), extractNumbersRegex.FindAllString(splitLine[1], -1)
}

func calculatePoints(winNumbers, numbers []string) int64 {
	var points int64
	for _, winNumber := range winNumbers {
		if slices.Contains(numbers, winNumber) {
			if points == 0 {
				points++

				continue
			}

			points *= 2
		}
	}

	return points
}
