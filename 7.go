package main

import "fmt"
import "time"

func main() {
	fmt.Println(nthPrime(6))
	fmt.Println(nthPrime(10001))
}

func nthPrime(n int) int {
	var res []int
	prime := 1
	for len(res) < n {
		prime = nextPrimeNumber(prime)
		res = append(res, prime)
		println(len(res))
	}
	return res[len(res)-1]
}

func isPrimeNumber(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}
	if n%5 == 0 {
		return false
	}
	if n%7 == 0 {
		return false
	}
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrimeNumber(n int) int {
	i := n + 1
	for {
		if isPrimeNumber(i) {
			return i
		}
		i++
	}
}
