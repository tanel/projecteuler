package main

import "io/ioutil"
import "fmt"
import "strings"
import "sort"

func main() {
	b, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Replace(string(b), "\"", "", -1)
	names := strings.Split(s, ",")

	sort.Strings(names)
	fmt.Println(names)
}
