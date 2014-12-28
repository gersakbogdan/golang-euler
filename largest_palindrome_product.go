package main

import "fmt"

func rev(n uint64) uint64 {
	var reverse uint64
	reverse = 0

	for n > 0 {
		reverse = 10 * reverse + n % 10
		n /= 10
	}

	return reverse
}

func isPalindrome(n uint64) bool {
	return n == rev(n)
}

func largestPalindrome() uint64 {
	var a, b, lp uint64
	a = 999
	lp = 0

	for a >= 100 {
		b = 999
		for b >= a {
			if (a * b) < lp {
				break;
			}
			if (isPalindrome(a*b) && a*b > lp) {
				lp = a * b
			}
			b--
		}
		a--
	}

	return lp
}

func main() {
	fmt.Println(largestPalindrome())
}