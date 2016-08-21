package main

import "fmt"

func main() {
	fmt.Println(d(220)) // = 284
	fmt.Println(d(284)) // = 200

	fmt.Println(sumOfAmicableNumbers(10000))
}

func sumOfAmicableNumbers(max int) int {
	res := 0
	for i := 1; i < max; i++ {
		a := i
		b := d(a)
		if d(b) == a && a != b {
			res += i
		}
	}
	return res
}

// the sum of proper divisors of n (numbers less than n which divide evenly into n).
func d(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			res += i
		}
	}
	return res
}
