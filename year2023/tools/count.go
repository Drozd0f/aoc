package tools

func Count[T comparable](s []T, e T) int {
	var count int
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == e {
			count++
		}
	}

	return count
}
