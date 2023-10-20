/* Take N as input, if multiple of 3 print fizz,
if multiple of 5 print buzz, if multiple of 3 & 5 print fizzbuzz
https://leetcode.com/problems/fizz-buzz/
*/

package main

import (
	"fmt"
	"strconv"
)

func checkNum(n int) []string {
	out := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out = append(out, "FizzBuzz")
		} else if i%3 == 0 {
			out = append(out, "Fizz")
		} else if i%5 == 0 {
			out = append(out, "Buzz")
		} else {
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}

func main() {
	n := 20
	out := checkNum(n)
	fmt.Println(out)
}
