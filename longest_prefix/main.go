package main

import (
	"fmt"
	"strings"
)

func longestPrefix(s []string) {
	pref := s[0]

	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], pref) {
			fmt.Println(s[i], pref)
			pref = pref[:len(pref)-1]
			fmt.Println(pref)
		}
	}
	fmt.Println(pref)
}

func main() {
	s := []string{"flower", "flow", "flight"}
	longestPrefix(s)
}
