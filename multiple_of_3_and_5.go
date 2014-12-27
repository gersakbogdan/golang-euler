package main

import "fmt"

var target int = 999

func sumMultiple(nr int) int {
	n := target / nr
	return nr * ((n * (n + 1)) / 2)
}

func result() int {
	return sumMultiple(3) + sumMultiple(5) - sumMultiple(15)
}

func main() {
	fmt.Println(result())
}
