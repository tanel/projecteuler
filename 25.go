package main

import "fmt"
import "strconv"

func main() {
	i := 1
	for {
		i *= 10
		n := fib(i)
		s := strconv.Itoa(n)
		fmt.Println(i, len(s))
		if len(s) == 50 {
			fmt.Println("result", i)
			break
		}
	}
}

var cache = make(map[int]int)

// https://en.wikipedia.org/wiki/Fibonacci_number
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	a, exists := cache[n-1]
	if !exists {
		a = fib(n - 1)
		cache[n-1] = a
	}
	b, exists := cache[n-2]
	if !exists {
		b = fib(n - 2)
		cache[n-2] = b
	}
	return a + b
}
