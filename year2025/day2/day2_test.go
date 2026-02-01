package day2

import (
	"testing"

	"github.com/Drozd0f/reader"
	"github.com/Drozd0f/year2025/day2/input"
)

var parsedInput [][2]int64

func init() {
	parsedInput = Parse(
		reader.ReadAllFile(input.Input, "input.txt"),
	)
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day2
// cpu: Apple M3 Max
// BenchmarkPart1String
// BenchmarkPart1String-14	60	19679385 ns/op	10832618 B/op	1103544 allocs/op
func BenchmarkPart1String(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		answer := Part1String(parsedInput)
		_ = answer
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day2
// cpu: Apple M3 Max
// BenchmarkPart1Int
// BenchmarkPart1Int-14	85	13841550 ns/op	0 B/op	0 allocs/op
func BenchmarkPart1Int(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		answer := Part1Int(parsedInput)
		_ = answer
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day2
// cpu: Apple M3 Max
// BenchmarkPart2String
// BenchmarkPart2String-14	13	88010026 ns/op	17259576 B/op	1822166 allocs/op
func BenchmarkPart2String(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		answer := Part2String(parsedInput)
		_ = answer
	}
}

// goos: darwin
// goarch: arm64
// pkg: github.com/Drozd0f/year2025/day2
// cpu: Apple M3 Max
// BenchmarkPart2Int
// BenchmarkPart2Int-14	2	963024354 ns/op	8 B/op	0 allocs/op
func BenchmarkPart2Int(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		answer := Part2Int(parsedInput)
		_ = answer
	}
}
