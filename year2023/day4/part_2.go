package day4

import (
	"regexp"
	"slices"
	"strconv"
)

func CalculatePart2(fileName string) int64 {
	lines := ReadAllFile(fileName)

	cardsMap := make(map[int]int, len(lines))

	var result int64
	for _, line := range lines {
		cardNumber := extractCardNumber(line)
		clonesCnt := calculateClones(extractNumbers(line))
		cardsMap[cardNumber]++

		for cloneNumber := 1; cloneNumber <= clonesCnt; cloneNumber++ {
			cardsMap[cardNumber+cloneNumber] += cardsMap[cardNumber]
		}
	}

	for _, v := range cardsMap {
		result += int64(v)
	}

	return result
}

var extractCardNumberRegex = regexp.MustCompile(`Card\s+(\d+):`)

func extractCardNumber(line string) int {
	number, err := strconv.Atoi(extractCardNumberRegex.FindStringSubmatch(line)[1])
	if err != nil {
		panic(err)
	}

	return number
}

func calculateClones(winNumbers, numbers []string) int {
	var clonesCnt int
	for _, winNumber := range winNumbers {
		if slices.Contains(numbers, winNumber) {
			clonesCnt++
		}
	}

	return clonesCnt
}
