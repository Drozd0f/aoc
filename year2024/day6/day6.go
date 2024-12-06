package day6

import "fmt"

type position [2]int

const (
	guardCh    = "^"
	obstructCh = "#"
)

func parseInput(lines []string) (position, [][]string) {
	var guardPos position
	field := make([][]string, len(lines))

	for idx, line := range lines {
		parsedLine := make([]string, len(line))
		for jdx, ch := range line {
			if string(ch) == guardCh {
				guardPos = position{idx, jdx}
			}

			parsedLine[jdx] = string(ch)
		}

		field[idx] = parsedLine
	}

	return guardPos, field
}

func (p *position) toString() string {
	return fmt.Sprintf("y=%d;x=%d", p[0], p[1])
}
