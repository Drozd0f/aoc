package day8

import (
	"regexp"
)

const (
	rootNodeName = "AAA"
	exitNodeName = "ZZZ"

	rootNodeSymbol = "A"
	exitNodeSymbol = "Z"

	leftNode  = "L"
	rightNode = "R"
)

type nodesMap map[string]map[string]string

var nodeRegex = regexp.MustCompile(`\w+`)

func newNodes(lines []string) nodesMap {
	nodes := make(nodesMap, len(lines))
	for _, line := range lines {
		n := nodeRegex.FindAllString(line, -1)
		nodes[n[0]] = map[string]string{
			leftNode:  n[1],
			rightNode: n[2],
		}
	}

	return nodes
}
