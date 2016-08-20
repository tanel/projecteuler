package main

import "fmt"

func main() {
	fmt.Println(sumOfSquares(10))
	fmt.Println(squareOfSum(10))
	fmt.Println(difference(10))

	fmt.Println(sumOfSquares(100))
	fmt.Println(squareOfSum(100))
	fmt.Println(difference(100))
}

func difference(max int) int {
	return squareOfSum(max) - sumOfSquares(max)
}

func sumOfSquares(max int) int {
	res := 0
	for i := 1; i <= max; i++ {
		res += i * i
	}
	return res
}

func squareOfSum(max int) int {
	res := 0
	for i := 1; i <= max; i++ {
		res += i
	}
	return res * res
}
