package main

import "fmt"

func main() {
	fmt.Println(isPythagoreanTriplet(3, 4, 5))

	const limit = 1000
	for a := 0; a < limit; a++ {
		for b := 1; b < limit; b++ {
			for c := 2; c < limit; c++ {
				if a >= b {
					continue
				}
				if b >= c {
					continue
				}
				if isPythagoreanTriplet(a, b, c) {
					sum := a + b + c
					if sum == 1000 {
						product := a * b * c
						fmt.Println(product)
					}
				}
			}
		}
	}
}

func isPythagoreanTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}
