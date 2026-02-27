package day_number

import (
	"testing"
)

// BenchmarkPageIsValid-12    	11834248	       101.9 ns/op
func Benchmark1(b *testing.B) {
	for range b.N {
	}
}

// BenchmarkGetMiddleNumber-12    	 4672236	       253.4 ns/op
func Benchmark2(b *testing.B) {
	for range b.N {
	}
}
