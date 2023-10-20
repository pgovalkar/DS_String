package main

import "fmt"

func reverseString(s []byte) {
	// if len(s) == 0 {
	// 	fmt.Println("empty string")
	// }
	// n := len(s)

	// for i := 0; i < n/2; i++ {
	// 	s[i], s[n-i-1] = s[n-i-1], s[i]
	// }
	// fmt.Println(string(s))

	j := len(s) - 1
	fmt.Println(j)
	for i := 0; i < len(s)/2; i++ {
		fmt.Println(i, j)
		s[i], s[j] = s[j], s[i]
		j--
	}
	fmt.Println(string(s))
}

func main() {
	//s := "qwertyiii"
	s := []byte{'H', 'a', 'n', 'n', 'a'}
	reverseString(s)
}
