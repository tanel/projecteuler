package main

import "fmt"

func main() {
	fmt.Println(divisors(28))

	// Find a triangle number with over 500 divisors
	d := 0
	i := 1
	n := 0
	for d <= 500 {
		n += i
		i++
		d = divisors(n)
		fmt.Println(n, d)
	}
}

func divisors(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res++
		}
	}
	return res
}
