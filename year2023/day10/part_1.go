package day10

import (
	"fmt"
	"slices"
)

func parceSquareLoop(lines []string) ([][]string, []int64) {
	var startingPipeIdx []int64

	squareLoop := make([][]string, len(lines))
	for loopIdx, line := range lines {
		tiles := make([]string, len(lines[0]))
		for tileIdx, tile := range line {
			tiles[tileIdx] = string(tile)
			if string(tile) == startingPipe {
				startingPipeIdx = []int64{int64(loopIdx), int64(tileIdx)}
			}
		}

		squareLoop[loopIdx] = tiles
	}

	return squareLoop, startingPipeIdx
}

func CalculatePartOne(lines []string) {
	fmt.Println(calculatePartOne(parceSquareLoop(lines)))
}

func calculatePartOne(squareLoop [][]string, startPipeIdx []int64) int64 {
	lenSquareLoop := int64(len(squareLoop))

	validMoves := make([][]int64, 0, len(moveMap))
	var verticalMove, horizontalMove int64
	for moveName, moveIdx := range moveMap {
		verticalMove = startPipeIdx[0] - moveIdx[0]
		if verticalMove < 0 || verticalMove > lenSquareLoop {
			continue
		}

		horizontalMove = startPipeIdx[1] + moveIdx[1]
		if horizontalMove < 0 || horizontalMove > lenSquareLoop {
			continue
		}

		if slices.Contains(validMovePipes[moveName], squareLoop[verticalMove][horizontalMove]) {
			validMoves = append(validMoves, []int64{verticalMove, horizontalMove})
		}
	}

	results := make([]int64, len(validMoves))
	for idx, validMove := range validMoves {
		results[idx] = searchPipeRoute(
			int64(len(squareLoop)),
			squareLoop,
			[]int64{validMove[0], validMove[1]},
			startPipeIdx,
		)
	}

	return slices.Max(results) / 2
}

func searchPipeRoute(size int64, squareLoop [][]string, curPipeIdx, prevPipeIdx []int64) int64 {
	if slices.Contains(curPipeIdx, -1) || slices.Contains(curPipeIdx, size) {
		return 0
	}

	if slices.Equal(curPipeIdx, prevPipeIdx) {
		return 0
	}

	pipe := squareLoop[curPipeIdx[0]][curPipeIdx[1]]
	switch pipe {
	case verticalPipe:
		verticalMove := curPipeIdx[0] + curPipeIdx[0] - prevPipeIdx[0]

		return 1 + searchPipeRoute(size, squareLoop, []int64{verticalMove, curPipeIdx[1]}, curPipeIdx)
	case ewPipe:
		horizontalMove := curPipeIdx[1] + curPipeIdx[1] - prevPipeIdx[1]

		return 1 + searchPipeRoute(size, squareLoop, []int64{curPipeIdx[0], horizontalMove}, curPipeIdx)
	case nePipe:
		verticalMove := curPipeIdx[0]
		horizontalMove := curPipeIdx[1]

		if curPipeIdx[1] == prevPipeIdx[1] {
			horizontalMove += 1
		}

		if curPipeIdx[0] == prevPipeIdx[0] {
			verticalMove -= 1
		}

		return 1 + searchPipeRoute(size, squareLoop, []int64{verticalMove, horizontalMove}, curPipeIdx)
	case nwPipe:
		verticalMove := curPipeIdx[0]
		horizontalMove := curPipeIdx[1]

		if curPipeIdx[1] == prevPipeIdx[1] {
			horizontalMove -= 1
		}

		if curPipeIdx[0] == prevPipeIdx[0] {
			verticalMove -= 1
		}

		return 1 + searchPipeRoute(size, squareLoop, []int64{verticalMove, horizontalMove}, curPipeIdx)
	case swPipe:
		verticalMove := curPipeIdx[0]
		horizontalMove := curPipeIdx[1]

		if curPipeIdx[0] == prevPipeIdx[0] {
			verticalMove += 1
		}

		if curPipeIdx[1] == prevPipeIdx[1] {
			horizontalMove -= 1
		}

		return 1 + searchPipeRoute(size, squareLoop, []int64{verticalMove, horizontalMove}, curPipeIdx)
	case sePipe:
		verticalMove := curPipeIdx[0]
		horizontalMove := curPipeIdx[1]

		if curPipeIdx[0] == prevPipeIdx[0] {
			verticalMove += 1
		}

		if curPipeIdx[1] == prevPipeIdx[1] {
			horizontalMove += 1
		}

		return 1 + searchPipeRoute(size, squareLoop, []int64{verticalMove, horizontalMove}, curPipeIdx)
	case ground:
		return 0
	case startingPipe:
		return 1
	}

	return 0
}

//fmt.Println(prevPipeIdx, squareLoop[prevPipeIdx[0]][prevPipeIdx[1]])
//fmt.Println(curPipeIdx, squareLoop[curPipeIdx[0]][curPipeIdx[1]])
//fmt.Println(verticalMove, horizontalMove)
