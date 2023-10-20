//https://leetcode.com/problems/excel-sheet-column-number/

package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	out := 0
	pow := 1
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		out = out + int(s[i]-'A'+1)*pow
		pow = pow * 26
	}

	return out
}

func main() {
	s := "AAA"
	out := titleToNumber(s)
	fmt.Println(out)
}
