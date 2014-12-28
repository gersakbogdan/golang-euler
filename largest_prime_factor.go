package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n uint64) uint64 {
	var factor, lf uint64

	factor = 2
	lf = 1
	maxFactor := uint64(math.Sqrt(float64(n))) 

	for n % factor == 0 {
		n /= factor
		lf = 2
	}

	factor++

	for n > 1 && factor <= maxFactor {
		for n % factor == 0 {
			n /= factor
		}
		lf = factor
		factor += 2
	}

	return lf
}

func main() {
	var n uint64
	fmt.Scan(&n)

	factor := largestPrimeFactor(n)
	fmt.Println(factor)
}