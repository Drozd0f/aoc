package day4

import (
	"github.com/Drozd0f/tools"
)

const mAS = "MAS"

func SolutionForPartTwo(lines []string) int {
	parsedLines := parseInput(lines)

	centerChar := string(mAS[1])

	var result int
	for idx, line := range parsedLines {
		for jdx, char := range line {
			if char == centerChar {
				result += sumXMAS(parsedLines, [2]int{idx, jdx})
			}
		}
	}

	return result
}

func sumXMAS(lines [][]string, centerCharIdx [2]int) int {
	searchDown := len(lines[centerCharIdx[0]:]) >= 2
	searchLeft := len(lines[centerCharIdx[0]][:centerCharIdx[1]+1]) >= 2
	searchUp := len(lines[:centerCharIdx[0]+1]) >= 2
	searchRight := len(lines[centerCharIdx[0]][centerCharIdx[1]:]) >= 2

	if !(searchDown && searchLeft && searchUp && searchRight) {
		return 0
	}

	var (
		diagonal string
		result   int
	)

	// check first diagonal
	diagonal = lines[centerCharIdx[0]-1][centerCharIdx[1]-1] +
		lines[centerCharIdx[0]][centerCharIdx[1]] +
		lines[centerCharIdx[0]+1][centerCharIdx[1]+1]
	if diagonal == mAS || tools.Reverse(diagonal) == mAS {
		result++
	}

	// check second diagonal
	diagonal = lines[centerCharIdx[0]-1][centerCharIdx[1]+1] +
		lines[centerCharIdx[0]][centerCharIdx[1]] +
		lines[centerCharIdx[0]+1][centerCharIdx[1]-1]
	if diagonal == mAS || tools.Reverse(diagonal) == mAS {
		result++
	}

	return result / 2
}
