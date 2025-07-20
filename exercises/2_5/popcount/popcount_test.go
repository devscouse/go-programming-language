package popcount

import "testing"

func BenchmarkPopCountTableLookup(b *testing.B) {
	for b.Loop() {
		PopCountTableLookup(10)
	}
}

func BenchmarkPopCountBitClear(b *testing.B) {
	for b.Loop() {
		PopCountBitClear(10)
	}
}
