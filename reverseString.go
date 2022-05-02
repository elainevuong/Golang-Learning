/*
Complete the solution so that it reverses all of the words within the string passed in.

Example:

"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"
*/

package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	wordSlice := strings.Split(str, " ")
	var resultSlice []string

	for i := len(wordSlice) - 1; i >= 0; i-- {
		resultSlice = append(resultSlice, wordSlice[i])

	}

	return strings.Join(resultSlice, " ")
}

func main() {
	fmt.Println(ReverseWords("The greatest victory is that which requires no battle"))
}
