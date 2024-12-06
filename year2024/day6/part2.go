package day6

import (
	"fmt"

	"github.com/Drozd0f/tools"
)

var ban position

func SolutionForPartTwo(lines []string) int {
	guardPos, fields := parseInput(lines)

	ban = position{guardPos[0], guardPos[1]}

	var angle int
	positions := map[string]bool{}
	for {
		if moveV2(&guardPos, fields, angle, positions) {
			break
		}

		angle = (angle + 90) % 360
	}

	return len(positions)
}

func moveV2(guardPos *position, fields [][]string, angle int, positions map[string]bool) bool {
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
		if !tools.Between(guardPos[0]+movePos[0], 0, yLimit) ||
			!tools.Between(guardPos[1]+movePos[1], 0, xLimit) {
			return true
		}

		nextCh = fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]]
		if nextCh == obstructCh {
			return false
		}

		if !(ban[0] == guardPos[0]+movePos[0] && ban[1] == guardPos[1]+movePos[1]) {
			tmp := fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]]
			fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]] = obstructCh

			if isCircle(*guardPos, fields, (angle+90)%360) {
				positions[fmt.Sprintf("y=%d;x=%d", guardPos[0]+movePos[0], guardPos[1]+movePos[1])] = true
				fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]] = "O"
			} else {
				fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]] = tmp
			}
		}

		guardPos[0] += movePos[0]
		guardPos[1] += movePos[1]
	}
}

func isCircle(guardPos position, fields [][]string, angle int) bool {
	yLimit := len(fields) - 1
	xLimit := len(fields[0]) - 1
	movePos := changePos(angle)

	var nextCh string
	localPos := map[string]bool{}
	for {
		localPos[fmt.Sprintf("%s:%d", guardPos.toString(), angle)] = true

		if !tools.Between(guardPos[0]+movePos[0], 0, yLimit) ||
			!tools.Between(guardPos[1]+movePos[1], 0, xLimit) {
			return false
		}

		nextCh = fields[guardPos[0]+movePos[0]][guardPos[1]+movePos[1]]
		if nextCh == obstructCh {
			angle = (angle + 90) % 360
			movePos = changePos(angle)
		}

		guardPos[0] += movePos[0]
		guardPos[1] += movePos[1]
		if localPos[fmt.Sprintf("%s:%d", guardPos.toString(), angle)] {
			return true
		}
	}
}

func changePos(angle int) position {
	switch angle {
	case 0:
		return moveUp
	case 90:
		return moveRight
	case 180:
		return moveDown
	case 270:
		return moveLeft
	}

	return moveUp
}
