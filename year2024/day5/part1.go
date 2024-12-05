package day5

import (
	"slices"

	"github.com/Drozd0f/tools"
)

func SolutionForPartOne(lines []string) int {
	rules, pages := parseInput(lines)

	var result int
	for _, page := range pages {
		var (
			middleNumber string
			pageIsValid  bool
		)
		middleIdx := len(page) / 2
		for idx, number := range page {
			if idx == middleIdx {
				middleNumber = number
			}

			if !upRulesIsValid(rules, number, page[:idx]) || !downRulesIsValid(rules, number, page[idx+1:]) {
				break
			}

			if idx == len(page)-1 {
				pageIsValid = true
			}
		}

		if pageIsValid {
			result += tools.MustIntFromString(middleNumber)
		}
	}

	return result
}

func upRulesIsValid(rules Rules, number string, page []string) bool {
	downRule := rules[number][down]
	for _, upNumber := range page {
		if slices.Contains(downRule, upNumber) {
			return false
		}
	}

	return true
}

func downRulesIsValid(rules Rules, number string, page []string) bool {
	upRule := rules[number][up]
	for _, downNumber := range page {
		if slices.Contains(upRule, downNumber) {
			return false
		}
	}

	return true
}

func pageIsValidTest(rules Rules, page []string) bool {
	for idx, number := range page {
		if !upRulesIsValid(rules, number, page[:idx]) || !downRulesIsValid(rules, number, page[idx+1:]) {
			break
		}

		if idx == len(page)-1 {
			return true
		}
	}

	return false
}
