package main

import "fmt"

var abundant = make(map[int]bool)

func main() {
	fmt.Println(sumOfProperDivisors(28))
	fmt.Println(isPerfectNumber(28))
	fmt.Println(isAbundantNumber(12))

	const limit = 28123

	// find abundant numbers
	for i := 1; i < limit; i++ {
		if isAbundantNumber(i) {
			abundant[i] = true
		}
	}

	// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
	res := 0
	for i := 1; i < limit; i++ {
		if !isSumOfTwoAbundantNumbers(i) {
			res += i
		}
	}
	fmt.Println(res)
}

func isSumOfTwoAbundantNumbers(n int) bool {
	for k := range abundant {
		a := n - k
		if a < 12 {
			continue
		}
		if _, isAbundant := abundant[a]; isAbundant {
			return true
		}
	}
	return false
}

func sumOfProperDivisors(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			res += i
		}
	}
	return res
}

func isPerfectNumber(n int) bool {
	return sumOfProperDivisors(n) == n
}

func isAbundantNumber(n int) bool {
	return sumOfProperDivisors(n) > n
}
