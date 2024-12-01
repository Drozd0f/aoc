package day1

const part1InputSeparator = "   "

func SolutionForPart1(lines []string) int64 {
	left, right := parseInput(lines)
	left, right = quickSort(left), quickSort(right)
	var result int64
	for idx := 0; idx < len(left); idx++ {
		result += abs(left[idx], right[idx])
	}

	return result
}

func quickSort(slice []int64) []int64 {
	if len(slice) <= 1 {
		return slice
	}

	result := make([]int64, 0, len(slice))
	pivot := slice[len(slice)/2]

	left := sortLeft(slice, pivot)
	middle := sortMiddle(slice, pivot)
	right := sortRight(slice, pivot)

	result = append(result, quickSort(left)...)
	result = append(result, middle...)
	result = append(result, quickSort(right)...)

	return result
}

func sortLeft(slice []int64, pivot int64) []int64 {
	result := make([]int64, 0, len(slice))
	for _, val := range slice {
		if pivot > val {
			result = append(result, val)
		}
	}

	return result
}

func sortMiddle(slice []int64, pivot int64) []int64 {
	result := make([]int64, 0, len(slice))
	for _, val := range slice {
		if pivot == val {
			result = append(result, val)
		}
	}

	return result
}

func sortRight(slice []int64, pivot int64) []int64 {
	result := make([]int64, 0, len(slice))
	for _, val := range slice {
		if pivot < val {
			result = append(result, val)
		}
	}

	return result
}

func abs(n1, n2 int64) int64 {
	if n1 > n2 {
		return n1 - n2
	}

	return n2 - n1
}
