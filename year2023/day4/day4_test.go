package day4

import "testing"

func BenchmarkDayFourPartOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculatePart1("input.txt")
	}
}

func BenchmarkDayFourPartTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculatePart2("input.txt")
	}
}
