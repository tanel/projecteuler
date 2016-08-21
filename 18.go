package main

import "fmt"
import "strings"
import "strconv"

func main() {
	vertices := parse(triangle2)
	tip := parseGraph(vertices)

	findTotals(tip)

	fmt.Println(tip.largestTotal())
}

func findTotals(v *vertex) {
	if v == nil {
		return
	}
	if v.left != nil {
		findTotals(v.left)
	}
	if v.right != nil {
		findTotals(v.right)
	}
	left := v.digit
	if v.left != nil {
		left += v.left.largestTotal()
	}
	right := v.digit
	if v.right != nil {
		right += v.right.largestTotal()
	}
	v.totalLeft = left
	v.totalRight = right
}

const triangle = `
3
7 4
2 4 6
8 5 9 3`

const triangle2 = `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func parseGraph(input [][]*vertex) *vertex {
	var tip *vertex
	for i, level := range input {
		for j, v := range level {
			if i+1 < len(input) {
				v.left = input[i+1][j]
				v.right = input[i+1][j+1]
			}
			if nil == tip {
				tip = v
			}
		}
	}
	return tip
}

type vertex struct {
	left       *vertex
	right      *vertex
	digit      int
	totalLeft  int
	totalRight int
}

func (v *vertex) largestTotal() int {
	if v.totalLeft > v.totalRight {
		return v.totalLeft
	}
	return v.totalRight
}

func (v *vertex) String() string {
	res := strconv.Itoa(v.digit)
	if v.left != nil || v.right != nil {
		res += " -> "
		if v.left != nil {
			res += strconv.Itoa(v.left.digit)
		}
		if v.right != nil {
			res += ", " + strconv.Itoa(v.right.digit)
		}
		res += " (" + strconv.Itoa(v.totalLeft) + ", " + strconv.Itoa(v.totalRight) + ")"
	}
	return res
}

func parse(input string) [][]*vertex {
	var res [][]*vertex
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		var digits []*vertex
		for _, field := range fields {
			digit, _ := strconv.Atoi(field)
			var v vertex
			v.digit = digit
			digits = append(digits, &v)
		}
		res = append(res, digits)
	}
	return res
}
