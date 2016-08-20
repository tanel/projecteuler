package main

import "fmt"
import "math/big"
import "strconv"

func main() {
	n := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(15), nil)
	fmt.Println(n, sumOfDigits(n))

	n = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(n, sumOfDigits(n))
}

func sumOfDigits(n *big.Int) int {
	sum := 0
	for _, s := range n.String() {
		digit, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}
		sum += digit
	}
	return sum
}
