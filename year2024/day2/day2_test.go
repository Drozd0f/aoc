package day2

import (
	"testing"
)

// BenchmarkSolutionForPartOne-12    	 1884061	       622.3 ns/op
func BenchmarkSolutionForPartOne(b *testing.B) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	for range b.N {
		SolutionForPartOne(input)
	}
}

// BenchmarkSolutionForPartTwo-12    	 1437045	       825.0 ns/op
func BenchmarkSolutionForPartTwo(b *testing.B) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	for range b.N {
		SolutionForPartTwo(input)
	}
}
