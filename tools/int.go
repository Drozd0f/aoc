package tools

import (
	"log"
	"strconv"
)

func MustInt64FromString(s string) int64 {
	return int64(MustIntFromString(s))
}

func MustIntFromString(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("strconv atoi: %s", s)
	}

	return number
}

func Abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}

func Between(n, left, right int) bool {
	return left <= n && n <= right
}
