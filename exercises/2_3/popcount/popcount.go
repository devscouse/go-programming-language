/*
Package popcount returns the number of set bits.
Excercise 2.3: Rewrite PopCount to use a loop instead of a single expresison.
Compare the performance of the two versions.
*/
package popcount

// This is a precomputed population count table that contains the number of set
// bits for each possible 8-bit value
// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountSingleExpr(x uint64) int {
	return int(
		pc[byte(x>>(0))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var res byte
	for i := range 8 {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}
