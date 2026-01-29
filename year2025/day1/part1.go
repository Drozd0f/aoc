package day1

func Part1(tacts []int64) int64 {
	var result int64

	var pos int64 = 50
	for _, tact := range tacts {
		if _, pos = tactRotation(pos, tact); pos == 0 {
			result++
		}
	}

	return result
}
