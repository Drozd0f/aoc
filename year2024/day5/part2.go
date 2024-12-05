package day5

import (
	"slices"

	"github.com/Drozd0f/tools"
)

func SolutionForPartTwo(lines []string) int {
	rules, pages := parseInput(lines)

	var result int
	for _, page := range pages {
		result += getMiddleNumber(rules, page)
	}

	return result
}

func getDownIdx(rules Rules, number string, page []string) int {
	upRule := rules[number][up]

	var newIdx int
	for _, downNumber := range page {
		if slices.Contains(upRule, downNumber) {
			newIdx++
		}
	}

	return newIdx
}

func getMiddleNumber(rules Rules, page []string) int {
	var middleNumber string

	pageIsValid := true
	middleIdx := len(page) / 2
	for idx, number := range page {
		downIdx := getDownIdx(rules, number, page)

		if downIdx == middleIdx {
			middleNumber = number
		}

		if downIdx != idx {
			pageIsValid = false
		}
	}

	if !pageIsValid {
		return tools.MustIntFromString(middleNumber)
	}

	return 0
}
