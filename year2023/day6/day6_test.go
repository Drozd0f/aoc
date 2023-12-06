package day6

import (
	"testing"
)

var testRace = race{
	duration: 40817772,
	distance: 219101213651089,
}

var testRaces = []race{
	{duration: 40, distance: 219},
	{duration: 81, distance: 1012},
	{duration: 77, distance: 1365},
	{duration: 72, distance: 1089},
}

func BenchmarkDayFourPartOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateWinningWays(testRaces)
	}
}

func BenchmarkDayFourPartTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateWinningWaysPart2(testRace)
	}
}
