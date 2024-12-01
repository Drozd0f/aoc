package day1

import (
	"testing"
)

// BenchmarkSolutionForPart1-12    	 1529718	       792.4 ns/op
func BenchmarkSolutionForPart1(b *testing.B) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	for range b.N {
		SolutionForPart1(input)
	}
}

// BenchmarkSolutionForPart2-12    	 2881292	       430.5 ns/op
func BenchmarkSolutionForPart2(b *testing.B) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	for range b.N {
		SolutionForPart2(input)
	}
}
