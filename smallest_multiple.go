package main

import "fmt"

func smallestMultiple(n uint64) uint64 {
	var sm uint64
	sm = 1

	for n > 0 {
		sm = lcm(sm, n)
		n--
	}

	return sm
}

func lcm(a, b uint64) uint64 {
	return a * b / gcd(a, b)
}

func gcd(a, b uint64) uint64 {
	c := a % b

	if c == 0 {
		return b
	}

	return gcd(b, c)
}


func main() {
	var n uint64
	fmt.Scan(&n)
	
	sm := smallestMultiple(n)
	fmt.Println(sm)
}