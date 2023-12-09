package day4

import (
	"github.com/Drozd0f/year2023/day4/input"
	"testing"

	"github.com/Drozd0f/year2023/reader"
)

var (
	benchmarkLines []string
)

func BenchmarkDayFourPartOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculatePart1(reader.ReadAllFile(input.Input, "input.txt"))
	}
}

func BenchmarkDayFourPartTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculatePart2(reader.ReadAllFile(input.Input, "input.txt"))
	}
}
