package day8

import (
	"github.com/Drozd0f/year2023/reader"
	"testing"
)

func newNavigatorAndNodes() (string, nodesMap) {
	lines := reader.ReadAllFile("input/input.txt")

	return lines[0], newNodes(lines[2:])
}

func BenchmarkDayEightPartOne(b *testing.B) {
	navigator, nodes := newNavigatorAndNodes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartOne(navigator, nodes)
	}
}

func BenchmarkDayEightPartTwo(b *testing.B) {
	navigator, nodes := newNavigatorAndNodes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculatePartTwo(navigator, nodes)
	}
}
