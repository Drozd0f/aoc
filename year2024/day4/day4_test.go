package day4

import (
	"testing"
)

// BenchmarkSolutionForPartOne-12    	 1884061	       622.3 ns/op
func BenchmarkSolutionForPartOne(b *testing.B) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	for range b.N {
		SolutionForPartOne(input)
	}
}

// BenchmarkSolutionForPartTwo-12    	 1437045	       825.0 ns/op
func BenchmarkSolutionForPartTwo(b *testing.B) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	for range b.N {
		SolutionForPartTwo(input)
	}
}
