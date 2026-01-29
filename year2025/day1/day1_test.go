package day1

import (
	"testing"

	"github.com/Drozd0f/reader"
	"github.com/Drozd0f/year2025/day1/input"
)

var (
	tacts []int64
)

func init() {
	tacts = Parse(
		reader.ReadAllFile(input.Input, "input.txt"),
	)
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day1
// cpu: Apple M3 Max
// BenchmarkPart1
// BenchmarkPart1-14	115717	10453 ns/op	0 B/op	0 allocs/op
func BenchmarkPart1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		result := Part1(tacts)
		_ = result
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day1
// cpu: Apple M3 Max
// BenchmarkPart2
// BenchmarkPart2-14	113080	10786 ns/op	0 B/op	0 allocs/op
func BenchmarkPart2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		result := Part2(tacts)
		_ = result
	}
}
