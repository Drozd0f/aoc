package day9

import (
	"github.com/Drozd0f/year2023/day9/input"
	"github.com/Drozd0f/year2023/reader"
	"testing"
)

func testSequence() [][]int64 {
	return parceSequence(reader.ReadAllFile(input.Input, "input.txt"))
}

func BenchmarkDayEightPartOne(b *testing.B) {
	sequence := testSequence()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartOne(sequence)
	}
}

func BenchmarkDayEightPartTwo(b *testing.B) {
	sequence := testSequence()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartTwo(sequence)
	}
}
