package day6

import (
	"fmt"

	"github.com/Drozd0f/tools"
)

const plusCh = "+"

func SolutionForPartTwo(lines []string) int {
	guardPos, fields := parseInput(lines)

	var angle int
	//plusPos := map[string]bool{}
	uniquePos := map[string]bool{}
	for {
		if moveV2(&guardPos, fields, angle, uniquePos) {
			break
		}

		angle = (angle + 90) % 360
	}

	return len(uniquePos)
}

func moveV2(guardPos *position, fields [][]string, angle int, uniquePos map[string]bool) bool {
	movePos := changeAngle(angle)

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
			delete(uniquePos, guardPos.toString())

			return false
		}

		guardPos[0] += movePos[0]
		guardPos[1] += movePos[1]

		if scanObstruct(*guardPos, angle+90, fields, uniquePos) {
			uniquePos[guardPos.toString()] = true
		}
	}
}

func scanObstruct(guardPos position, angle int, fields [][]string, uniquePos map[string]bool) bool {
	yLimit := len(fields) - 1
	xLimit := len(fields[0]) - 1

	movePos := changeAngle(angle)
	newGuardPos := position{guardPos[0] + movePos[0], guardPos[1] + movePos[1]}

	for i := 0; i < 100; i++ {
		if !tools.Between(newGuardPos[0]+movePos[0], 0, yLimit) ||
			!tools.Between(newGuardPos[1]+movePos[1], 0, xLimit) {
			break
		}

		if newGuardPos[0] == guardPos[0] && newGuardPos[1] == guardPos[1] {
			uniquePos[guardPos.toString()] = true
			fmt.Println(len(uniquePos))

			return true
		}

		nextCh := fields[newGuardPos[0]+movePos[0]][newGuardPos[1]+movePos[1]]
		if nextCh == obstructCh {
			angle = (angle + 90) % 360
			movePos = changeAngle(angle)
		}

		newGuardPos[0] += movePos[0]
		newGuardPos[1] += movePos[1]
	}

	return false
}

func changeAngle(angle int) position {
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
