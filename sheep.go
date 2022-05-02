/*
If you can't sleep, just count sheep!!

Task:
Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.
*/

package main

import (
	"fmt"
	"strconv"
)

func countSheep(num int) string {
	var result string

	for i := 1; i <= num; i++ {
		strNum := strconv.Itoa(i)
		result += strNum + " sheep..."
	}

	return result
}

func main() {
	fmt.Println(countSheep(5))
	fmt.Println(countSheep(3))
}
