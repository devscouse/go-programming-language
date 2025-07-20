/*
Package popcount returns the number of set bits.
Excercise 2.4: Write a version of PopCount that counts bits by shifting its
argument through 64 bit positions, testing the rightmost bit each time. Compare
its performance with the table-lookup version.
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

func PopCountTableLookup(x uint64) int {
	return int(
		pc[byte(x>>(0))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountBitShift(x uint64) int {
	var count byte = 0
	for i := range 64 {
		count += byte((x >> i) & 1)
	}
	return int(count)
}
