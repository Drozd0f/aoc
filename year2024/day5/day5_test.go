package day5

import (
	"testing"
)

// BenchmarkPageIsValid-12    	11834248	       101.9 ns/op
func BenchmarkPageIsValid(b *testing.B) {
	rules := Rules{
		"47": {
			up:   []string{"75", "97", "61"},
			down: []string{"53", "13", "61", "29"},
		},
		"53": {
			up:   []string{"47", "75", "61", "97"},
			down: []string{"29", "13"},
		},
		"97": {
			up:   []string{},
			down: []string{"13", "61", "47", "29", "75", "53"},
		},
		"75": {
			up:   []string{"97"},
			down: []string{"29", "53", "47", "61", "13"},
		},
		"61": {
			up:   []string{"75", "97", "47"},
			down: []string{"13", "53", "29"},
		},
		"29": {
			up:   []string{"75", "97", "53", "61"},
			down: []string{"13"},
		},
		"13": {
			up:   []string{"47", "29", "75", "61", "97", "53"},
			down: []string{},
		},
	}
	page := []string{"97", "13", "75", "29", "47"}

	for range b.N {
		pageIsValidTest(rules, page)
	}
}

// BenchmarkGetMiddleNumber-12    	 4672236	       253.4 ns/op
func BenchmarkGetMiddleNumber(b *testing.B) {
	rules := Rules{
		"47": {
			up:   []string{"75", "97", "61"},
			down: []string{"53", "13", "61", "29"},
		},
		"53": {
			up:   []string{"47", "75", "61", "97"},
			down: []string{"29", "13"},
		},
		"97": {
			up:   []string{},
			down: []string{"13", "61", "47", "29", "75", "53"},
		},
		"75": {
			up:   []string{"97"},
			down: []string{"29", "53", "47", "61", "13"},
		},
		"61": {
			up:   []string{"75", "97", "47"},
			down: []string{"13", "53", "29"},
		},
		"29": {
			up:   []string{"75", "97", "53", "61"},
			down: []string{"13"},
		},
		"13": {
			up:   []string{"47", "29", "75", "61", "97", "53"},
			down: []string{},
		},
	}
	invalidPage := []string{"97", "13", "75", "29", "47"}

	for range b.N {
		getMiddleNumber(rules, invalidPage)
	}
}
