package main

import "fmt"

func main() {
	counts := 0
	for i := 1; i <= 1000; i++ {
		counts += letters(numberAsWord(i))
	}
	fmt.Println(counts)
}

func debug(n int) {
	fmt.Println(n, "=", numberAsWord(n))
}

var words = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var words2 = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func letters(input string) int {
	res := 0
	for _, s := range input {
		digit := string(s)
		if digit == "-" || digit == " " {
			continue
		}
		res++
	}
	return res
}

func numberAsWord(n int) string {
	if n == 0 {
		return ""
	} else if n > 0 && n <= 19 {
		return words[n-1]
	} else if n >= 20 && n <= 90 && n%10 == 0 {
		return words2[(n-20)/10]
	} else if n == 1000 {
		return "one thousand"
	} else if n >= 100 {
		hundreds := n / 100
		s := words[hundreds-1] + " hundred"
		if n-hundreds*100 > 0 {
			s += " and " + numberAsWord(n-hundreds*100)
		}
		return s
	}
	first := n / 10
	second := n % 10
	return words2[first-2] + "-" + words[second-1]
}
