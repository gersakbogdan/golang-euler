package main

import (
	"fmt"
)

func main() {
	var n uint64
	fmt.Scan(&n)

	sum := n * (n + 1) / 2
	sum_square := n * (n + 1) * (2*n + 1) / 6
	diff := sum*sum - sum_square

	fmt.Println(diff)
}
