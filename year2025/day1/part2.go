package day1

func Part2(tacts []int64) int64 {
	var result int64

	var pos int64 = 50
	for _, tact := range tacts {
		var rot int64
		rot, pos = tactRotation(pos, tact)
		if pos == 0 {
			result++
		}

		result += rot
	}

	return result
}
