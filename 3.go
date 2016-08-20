package main

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

import "fmt"

func main() {
	fmt.Println(findPrimeFactors(12))
	fmt.Println(findPrimeFactors(147))
	fmt.Println(findPrimeFactors(17))
	fmt.Println(findPrimeFactors(600851475143))
}

func findPrimeFactors(n int) []int {
	var res []int
	prime := nextPrimeNumber(1)
	for {
		if n%prime != 0 {
			prime = nextPrimeNumber(prime)
			continue
		}
		println(n, prime)
		res = append(res, prime)
		if prime == n {
			break
		}
		n = n / prime
		prime = nextPrimeNumber(1)
	}
	return res
}

func isPrimeNumber(n int) bool {
	if n < 2 {
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
