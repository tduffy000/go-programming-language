package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		b += pc[byte(x>>(uint8(i)*8))]
	}
	return int(b)
}

func PopCountShift(x uint64) int {
	pop := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			pop++
		}
	}
	return pop
}

func PopCountClear(x uint64) int {
	pop := 0
	for x != 0 {
		x = x & (x - 1)
		pop++
	}
	return pop
}
