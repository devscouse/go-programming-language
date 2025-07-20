package popcount

import "testing"

func BenchmarkPopCountTableLookup(b *testing.B) {
	for b.Loop() {
		PopCountTableLookup(10)
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for b.Loop() {
		PopCountBitShift(10)
	}
}
