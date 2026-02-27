package day6

import (
	"fmt"

	"github.com/Drozd0f/tools"
)

var (
	moveRight = position{0, 1}
	moveLeft  = position{0, -1}
	moveUp    = position{-1, 0}
	moveDown  = position{1, 0}
)

func SolutionForPartOne(lines []string) int {
	guardPos, fields := parseInput(lines)

	var angle int
	positions := map[string]bool{}
	for {
		if move(&guardPos, fields, angle, positions) {
			break
		}

		angle = (angle + 90) % 360
	}

	return len(positions)
}

func move(guardPos *position, fields [][]string, angle int, positions map[string]bool) bool {
	var movePos position
	switch angle {
	case 0:
		movePos = moveUp
	case 90:
		movePos = moveRight
	case 180:
		movePos = moveDown
	case 270:
		movePos = moveLeft
	}

	yLimit := len(fields) - 1
	xLimit := len(fields[0]) - 1

	var nextCh string
	for {
		positions[fmt.Sprintf("y=%d;x=%d", guardPos[0], guardPos[1])] = true

		if !tools.Between(guardPos[0]+movePos[0], 0, yLimit) ||
			!tools.Between(guardPos[1]+movePos[1], 0, xLimit) {
			return true
		}

		nextCh = fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]]
		if nextCh == obstructCh {
			return false
		}

		guardPos[0] += movePos[0]
		guardPos[1] += movePos[1]
	}
}
