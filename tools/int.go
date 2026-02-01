package tools

import (
	"log"
	"math"
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

func LenInt[T Int](d T) int {
	if Abs(d) < 10 {
		return 1
	}

	return int(math.Log10(float64(Abs(d)))) + 1
}

func CutFromInt[T Int](v T, start, end int) T {
	l := LenInt(v)
	if l == 1 {
		return v
	}

	if end == 0 {
		end = l
	}

	var res T
	for digit := 0; digit < end; digit++ {
		pow10 := T(math.Pow10(l - 1 - digit))
		if digit < start {
			v %= pow10

			res *= 10
			continue
		}

		number := v / pow10
		res = res*10 + number
		v -= number * pow10
	}

	return res
}

func CountFromInt[T Int](d, sub T) int {
	if sub == 0 {
		return 0
	}

	var (
		res    int
		window T
	)

	subLen := LenInt(sub)
	start, end := 0, subLen
	for end <= LenInt(d) {
		window = CutFromInt(d, start, end)
		if window == sub {
			start, end = end, end+subLen
			res++

			continue
		}

		start++
		end++
	}

	return res
}
