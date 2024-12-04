package day4

func parseInput(lines []string) [][]string {
	parsedLines := make([][]string, len(lines))
	for idx, line := range lines {
		parsedLine := make([]string, len(line))
		for jdx, char := range line {
			parsedLine[jdx] = string(char)
		}

		parsedLines[idx] = parsedLine
	}

	return parsedLines
}
