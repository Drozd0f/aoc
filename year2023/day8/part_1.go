package day8

func CalculatePartOne(lines []string) int64 {
	return calculatePartOne(lines[0], newNodes(lines[2:]))
}

func calculatePartOne(navigator string, nodes nodesMap) int64 {
	var isExitFound bool
	var nodeCnt int64

	curNode := rootNodeName
	for !isExitFound {
		for _, node := range navigator {
			if curNode == exitNodeName {
				isExitFound = true
				break
			}

			curNode = choiceNode(string(node), curNode, nodes)
			nodeCnt++
		}
	}

	return nodeCnt
}

func choiceNode(turn string, curNode string, nodes nodesMap) string {
	if string(turn) == leftNode {
		return nodes[curNode][leftNode]
	}

	return nodes[curNode][rightNode]
}
