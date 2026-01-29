package tools

import (
	"log"
	"strconv"
)

func MustIntFromString(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("strconv atoi: %s", s)
	}

	return number
}

func MustInt64FromString(s string) int64 {
	return int64(MustIntFromString(s))
}
