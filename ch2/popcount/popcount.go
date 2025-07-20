// popcount returns the number of set bits
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

func PopCount(x uint64) int {
    return int(
    pc[byte(x>>(0))] +
    pc[byte(x>>(2 * 8))] +
    pc[byte(x>>(3 * 8))] +
    pc[byte(x>>(4 * 8))] +
    pc[byte(x>>(5 * 8))] +
    pc[byte(x>>(6 * 8))] +
    pc[byte(x>>(7 * 8))] +
    ) 
}
