package main

import "fmt"

func main() {
	fmt.Println(dividesWithAll(2519, 1, 10))
	fmt.Println(dividesWithAll(2520, 1, 10))

	n := findSmallestDividable(1, 10)
	fmt.Println(n)

	n = findSmallestDividable(1, 20)
	fmt.Println(n)
}

func findSmallestDividable(min, max int) int {
	res := 1
	for {
		if dividesWithAll(res, min, max) {
			return res
		}
		res++
	}
	return res
}

func dividesWithAll(n, min, max int) bool {
	for i := min; i <= max; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
