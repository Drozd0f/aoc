package day2

import (
	"math"
	"strconv"
)

func Part1String(seqs [][2]int64) int64 {
	var res int64
	for _, seq := range seqs {
		for i := seq[0]; i <= seq[1]; i++ {
			s := strconv.Itoa(int(i))
			if len(s)%2 != 0 {
				i = int64(math.Pow10(len(s)))

				continue
			}

			if s[:len(s)/2] == s[len(s)/2:] {
				res += i
			}
		}
	}

	return res
}

func Part1Int(seqs [][2]int64) int64 {
	var res int64
	for _, seq := range seqs {
		for i := seq[0]; i <= seq[1]; i++ {
			digits := int(math.Log10(float64(i)))
			if digits%2 == 0 {
				i = int64(math.Pow10(digits + 1))

				continue
			}

			var left int64
			right := i
			for digit := int64(digits); digit > int64(digits)/2; digit-- {
				pow10 := int64(math.Pow10(int(digit)))
				if right < pow10 {
					left *= 10

					continue
				}

				v := right / pow10
				left = left*10 + v
				right -= v * pow10
			}

			if left == right {
				res += i
			}
		}
	}

	return res
}
