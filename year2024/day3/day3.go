package day3

import (
	"regexp"

	"github.com/Drozd0f/tools"
)

var (
	validMulRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
)

func calculateMul(line string) int64 {
	var res int64
	for _, sub := range validMulRe.FindAllStringSubmatch(line, -1) {
		res += mul(sub[1], sub[2])
	}

	return res
}

func mul(n1, n2 string) int64 {
	return tools.MustInt64FromString(n1) * tools.MustInt64FromString(n2)
}
