package main

import "strconv"
import "fmt"

func main() {
	fmt.Println(reverse("9009"))
	fmt.Println(isPalindrome(9009))

	n := findLargestPalindrome(100)
	fmt.Println(n)

	n = findLargestPalindrome(1000)
	fmt.Println(n)
}

func findLargestPalindrome(max int) int {
	largest := 0
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			product := i * j
			if product < largest {
				continue
			}
			if isPalindrome(product) {
				largest = product
			}
		}
	}
	return largest
}

func isPalindrome(n int) bool {
	a := strconv.Itoa(n)
	b := reverse(a)
	return a == b
}

func reverse(input string) string {
	res := ""
	for _, s := range input {
		res = string(s) + res
	}
	return res
}
