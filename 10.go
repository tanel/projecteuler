package main

import "fmt"

func main() {
	fmt.Println(sieveOfEratosthenes(10))
	fmt.Println(sumOfPrimes(10))
	fmt.Println(sumOfPrimes(2000000))
}

func sumOfPrimes(n int) int {
	primes := sieveOfEratosthenes(n)
	res := 0
	for _, n := range primes {
		res += n
	}
	return res
}

func sieveOfEratosthenes(N int) []int {
	var i, j int
	var a []bool
	for i := 0; i <= N+1; i++ {
		a = append(a, false)
	}

	for a[1], i = false, 2; i <= N; i++ {
		a[i] = true
	}

	for i = 2; i <= N/2; i++ {
		for j = 2; j <= N/i; j++ {
			a[i*j] = false
		}
	}

	var res []int
	for i := range a {
		if a[i] {
			res = append(res, i)
		}
	}

	return res
}
