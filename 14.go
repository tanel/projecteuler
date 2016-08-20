package main

import "fmt"

func main() {
	fmt.Println(collatzSequence(13))
	fmt.Println(collatzSequence(9))

	fmt.Println(longestChain(1000000))
}

func longestChain(max int) (int, int) {
	length := 0
	number := 0
	for i := 0; i < max; i++ {
		seq := collatzSequence(i)
		if len(seq) > length {
			length = len(seq)
			number = i
		}
	}
	return number, length
}

func collatzSequence(n int) []int {
	var res []int
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		res = append(res, n)
	}
	return res
}
