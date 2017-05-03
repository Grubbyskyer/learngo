package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var i uint64
	fmt.Println("using single exp:")
	start := time.Now()
	for i = 0; i < 1e8; i++ {
		PopCount1(i)
	}
	fmt.Printf("time cost: %.2fs\n", time.Since(start).Seconds())
	fmt.Println("using for-loop:")
	start = time.Now()
	for i = 0; i < 1e8; i++ {
		PopCount2(i)
	}
	fmt.Printf("time cost: %.2fs\n", time.Since(start).Seconds())
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var count byte
	var i uint
	for i = 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}
