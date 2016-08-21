package main

import "io/ioutil"
import "fmt"
import "strings"
import "sort"

func main() {
	names := parseNames()

	res := 0
	for i, name := range names {
		res += (i + 1) * worth(name)
	}
	fmt.Println(res)
}

func worth(name string) int {
	res := 0
	for _, s := range name {
		res += indexOf(string(s))
	}
	return res
}

func parseNames() []string {
	b, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Replace(string(b), "\"", "", -1)
	names := strings.Split(s, ",")

	sort.Strings(names)

	return names
}

func indexOf(input string) int {
	for i, s := range alphabet {
		if input == s {
			return i + 1
		}
	}
	panic("not found")
}

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
