package main

import "fmt"
import "math/big"
import "strconv"

func main() {
	f := factorial(10)
	fmt.Println(f)
	fmt.Println(sumOfDigits(f))

	f = factorial(100)
	fmt.Println(f)
	fmt.Println(sumOfDigits(f))
}

func factorial(n int64) *big.Int {
	if n == 1 {
		return big.NewInt(n)
	}
	res := big.NewInt(n)
	return res.Mul(res, factorial(n-1))
}

func sumOfDigits(n *big.Int) int {
	res := 0
	for _, digit := range n.String() {
		n, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		res += n
	}
	return res
}
