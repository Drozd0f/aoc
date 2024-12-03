package day3

import (
	"testing"
)

// BenchmarkSolutionForPartOne-12    	 1884061	       622.3 ns/op
func BenchmarkSolutionForPartOne(b *testing.B) {
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	for range b.N {
		SolutionForPartOne(input)
	}
}

// BenchmarkSolutionForPartTwo-12    	 1437045	       825.0 ns/op
func BenchmarkSolutionForPartTwo(b *testing.B) {
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	for range b.N {
		SolutionForPartTwo(input)
	}
}
