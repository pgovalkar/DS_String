package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	s = strings.Trim(s, " ")
	str := strings.SplitAfter(s, " ")
	lastword := str[len(str)-1]
	fmt.Println(lastword)
	out := len(lastword)
	return out
}

func main() {
	s := "   fly me   to   the moon  "
	out := lengthOfLastWord(s)
	fmt.Println(out)

}
