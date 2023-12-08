package day8

import (
	"fmt"
	"strings"
)

func CalculatePartTwo(lines []string) {
	fmt.Println(calculatePartTwo(lines[0], newNodes(lines[2:])))
}

func calculatePartTwo(navigator string, nodes nodesMap) int64 {
	curNodes := make([]string, 0, 8)
	for node := range nodes {
		if strings.Contains(node, rootNodeSymbol) {
			curNodes = append(curNodes, node)
		}
	}

	steps := make([]int64, len(curNodes))
	for idx, node := range curNodes {
		var step int64
		var isExitFound bool

		curNode := node
		for !isExitFound {
			for _, turn := range navigator {
				if strings.Contains(curNode, exitNodeSymbol) {
					isExitFound = true
					break
				}

				curNode = choiceNode(string(turn), curNode, nodes)
				step++
			}
		}

		steps[idx] = step
	}

	return findLCM(steps)
}

func updateCurNodes(turn string, nodes nodesMap, curNodes []string) ([]string, bool) {
	var isExitFound = true
	newCurNodes := make([]string, len(curNodes))
	for idx, curNode := range curNodes {
		if !strings.Contains(curNode, exitNodeSymbol) {
			isExitFound = false
		}

		newCurNodes[idx] = choiceNode(turn, curNode, nodes)
	}

	return newCurNodes, isExitFound
}
