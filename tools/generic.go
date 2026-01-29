package tools

type (
	Int interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Uint interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Float interface {
		~float32 | ~float64
	}
	Number interface {
		Int | Uint | Float
	}
	Ordered interface {
		Number | ~string
	}
)

func Between[T Ordered](e1, e2, v T) bool {
	return e1 < v && v < e2
}

func Count[T comparable](s []T, e T) int {
	var count int
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == e {
			count++
		}
	}

	return count
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}

	return a
}
