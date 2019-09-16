package main

import "fmt"

// GCD greatest common divisor  TAOCP 1.1E Euclid's Algorithm
func GCD(m, n int) int {
	var t int
	for n != 0 {
		t = n
		n = m % n
		m = t
	}
	return t
}

func main() {
	fmt.Println(GCD(544, 119))
}
