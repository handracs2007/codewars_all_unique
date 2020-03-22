package main

import "fmt"

func HasUniqueChar(str string) bool {
	var charMap = make(map[rune]int)

	for _, val := range str {
		var _, ok = charMap[val]

		if ok {
			return false
		} else {
			charMap[val] = 1
		}
	}

	return true
}

func main() {
	fmt.Println(HasUniqueChar("  nAa"))
	fmt.Println(HasUniqueChar("abcde"))
	fmt.Println(HasUniqueChar("++-"))
	fmt.Println(HasUniqueChar("AaBbC"))
}
