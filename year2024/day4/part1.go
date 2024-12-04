package day4

const xMAS = "XMAS"

var lenXMAS = len(xMAS)

func SolutionForPartOne(lines []string) int {
	parsedLines := parseInput(lines)

	firstChar := string(xMAS[0])

	var result int
	for idx, line := range parsedLines {
		for jdx, char := range line {
			if char == firstChar {
				result += sumWord(parsedLines, [2]int{idx, jdx})
			}
		}
	}

	return result
}

func sumWord(lines [][]string, firstCharIdx [2]int) int {
	searchDown := len(lines[firstCharIdx[0]:]) >= lenXMAS
	searchLeft := len(lines[firstCharIdx[0]][:firstCharIdx[1]+1]) >= lenXMAS
	searchUp := len(lines[:firstCharIdx[0]+1]) >= lenXMAS
	searchRight := len(lines[firstCharIdx[0]][firstCharIdx[1]:]) >= lenXMAS

	var result int
	result += sumHorizontalWord(lines[firstCharIdx[0]], firstCharIdx[1], searchLeft, searchRight)

	result += sumVerticalWord(lines, firstCharIdx, searchUp, searchDown)

	result += sumDiagonalWord(lines, firstCharIdx, searchLeft, searchRight, searchUp, searchDown)

	return result

}

func sumHorizontalWord(line []string, firstCharIdx int, left, right bool) int {
	var (
		leftWord, rightWord string
		result              int
	)

	// check left word
	if left {
		for _, char := range line[firstCharIdx-lenXMAS+1 : firstCharIdx+1] {
			leftWord = char + leftWord
		}

		if leftWord == xMAS {
			result++
		}
	}

	// check right word
	if right {
		for _, char := range line[firstCharIdx : firstCharIdx+lenXMAS] {
			rightWord += char
		}

		if rightWord == xMAS {
			result++
		}
	}

	return result
}

func sumVerticalWord(lines [][]string, firstCharIdx [2]int, up, down bool) int {
	var (
		downWord, upWord string
		result           int
	)

	// check up word
	if up {
		for _, char := range lines[firstCharIdx[0]-lenXMAS+1 : firstCharIdx[0]+1] {
			upWord = char[firstCharIdx[1]] + upWord
		}

		if upWord == xMAS {
			result++
		}
	}

	// check down word
	if down {
		for _, char := range lines[firstCharIdx[0] : firstCharIdx[0]+lenXMAS] {
			downWord += char[firstCharIdx[1]]
		}

		if downWord == xMAS {
			result++
		}
	}

	return result
}

func sumDiagonalWord(lines [][]string, firstCharIdx [2]int, left, right, up, down bool) int {
	var (
		diagonalLeft, diagonalRight string
		result                      int
	)

	// check down diagonal
	if down && (left || right) {
		for idx, line := range lines[firstCharIdx[0] : firstCharIdx[0]+lenXMAS] {
			if idx > lenXMAS-1 {
				break
			}

			if left {
				diagonalLeft += line[firstCharIdx[1]-idx]
			}

			if right {
				diagonalRight += line[firstCharIdx[1]+idx]
			}
		}

		if diagonalLeft == xMAS {
			result++
		}

		if diagonalRight == xMAS {
			result++
		}

		diagonalLeft, diagonalRight = "", ""
	}

	// check up diagonal
	if up && (left || right) {
		for idx, line := range lines[firstCharIdx[0]-lenXMAS+1 : firstCharIdx[0]+1] {
			if idx > firstCharIdx[0] {
				break
			}

			if left {
				diagonalLeft = line[firstCharIdx[1]-lenXMAS+1+idx] + diagonalLeft
			}

			if right {
				diagonalRight = line[firstCharIdx[1]+lenXMAS-1-idx] + diagonalRight
			}
		}

		if diagonalLeft == xMAS {
			result++
		}

		if diagonalRight == xMAS {
			result++
		}
	}

	return result
}
