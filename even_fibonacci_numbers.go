package main

import (
	"fmt"
	"time"
)

func fib_even_numbers(max int) int {
	a, b, sum := 1, 2, 2

	for {
		a, b = a + b + b, a + a + b + b + b

		if b >= max {
			return sum
		}

		sum += b
	}
}

func main() {
	start := time.Now()

	max := 4000000
	fmt.Println(fib_even_numbers(max))
	fmt.Println("Elapsed: ", time.Since(start))
}