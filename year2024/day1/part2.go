package day1

func SolutionForPart2(lines []string) int64 {
	left, right := parseInput(lines)
	rightMap := rightMapper(right)

	var result int64
	for _, val := range left {
		result += val * rightMap[val]
	}

	return result
}

func rightMapper(right []int64) map[int64]int64 {
	result := make(map[int64]int64, len(right))
	for _, val := range right {
		result[val] += 1
	}

	return result
}
