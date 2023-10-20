// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
package main

import (
	"fmt"
	"strconv"
)

func reversePolishNotation(in []string) int {
	out := []int{}

	var v1, v2, res int
	for i := 0; i < len(in); i++ {

		if in[i] == "+" || in[i] == "-" || in[i] == "/" || in[i] == "*" {
			l := len(out)
			fmt.Println("i:", in[i], "l: ", l)
			fmt.Println(out)
			v2 = out[l-1]
			v1 = out[l-2]
			fmt.Println("v1: ", v1, "v2: ", v2)
			out = out[:l-2]
		} else {
			fmt.Println(in[i])
			res, _ = strconv.Atoi(in[i])
		}

		if in[i] == "+" {
			res = v1 + v2
		} else if in[i] == "-" {
			res = v1 - v2
		} else if in[i] == "/" {
			res = v1 / v2
		} else if in[i] == "*" {
			res = v1 * v2
		}
		out = append(out, res)

	}
	out = append(out, res)
	reslen := len(out)
	return out[reslen-1]
}

func main() {
	s := []string{"4", "13", "5", "/", "+"}
	out := reversePolishNotation(s)
	fmt.Println(out)
}
