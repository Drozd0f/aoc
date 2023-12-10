package day9

import (
	"testing"
)

var testHistories = [][]int64{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
}

func BenchmarkDayEightPartOne(b *testing.B) {
	//sequence := testSequence()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartOne(testHistories)
	}
}

func BenchmarkDayEightPartTwo(b *testing.B) {
	//sequence := testSequence()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartTwo(testHistories)
	}
}
