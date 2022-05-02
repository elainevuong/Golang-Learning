/*
Write a function that checks if a given string
(case insensitive) is a palindrome.
*/

package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	if len(str) == 1 {
		return true
	}

	str = strings.ToLower(str)

	length := len(str)

	for i := 0; i < length; i++ {
		length--
		if str[i] != str[length] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsPalindrome("a"))
	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("Abba"))
	fmt.Println(IsPalindrome("hello"))
}
