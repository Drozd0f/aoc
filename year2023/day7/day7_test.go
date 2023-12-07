package day7

import (
	"testing"
)

var testPlayers = []player{
	{
		hand: "32T3K",
		bid:  765,
	},
	{
		hand: "T55J5",
		bid:  684,
	},
	{
		hand: "KK677",
		bid:  28,
	},
	{
		hand: "KTJJT",
		bid:  220,
	},
	{
		hand: "QQQJA",
		bid:  483,
	},
}

func BenchmarkDaySevenPartOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateWinners(testPlayers)
	}
}

func BenchmarkDaySevenPartTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateWinnersPartTwo(testPlayers)
	}
}
