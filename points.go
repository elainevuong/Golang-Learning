/*
Our football team finished the championship. The result of each match look like "x:y". Results of all matches are recorded in the collection.

For example: ["3:1", "2:2", "0:1", ...]

Write a function that takes such collection and counts the points of our team in the championship. Rules for counting points for each match:

if x > y: 3 points
if x < y: 0 point
if x = y: 1 point
Notes:

there are 10 matches in the championship
0 <= x <= 4
0 <= y <= 4
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	var points int
	for _, value := range games {
		result := strings.Split(value, ":")
		x, _ := strconv.Atoi(result[0])
		y, _ := strconv.Atoi(result[1])

		points += CalcPoints(x, y)
	}

	return points

}

func CalcPoints(x, y int) int {
	if x > y {
		return 3
	} else if x < y {
		return 0
	} else {
		return 1
	}
}

func main() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))

}
