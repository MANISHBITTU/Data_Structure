package main

import "fmt"

func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}

	if m%n == 0 {
		return n
	}

	return GCD(n, m%n)
}
func main() {
	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	fmt.Println("The GCD of", m, n, "is", GCD(m, n))
}
