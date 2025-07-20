package popcount

import "testing"

func BenchmarkPopCountSingleExpr(b *testing.B) {
	for b.Loop() {
		PopCountSingleExpr(10)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for b.Loop() {
		PopCountLoop(10)
	}
}
