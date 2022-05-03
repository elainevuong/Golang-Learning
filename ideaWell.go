/*
For every good kata idea there seem to be quite a few bad ones!

In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.
*/

package main

import "fmt"

func Well(ideas []string) string {
	ideaCount := make(map[string]int)

	for _, idea := range ideas {
		ideaCount[idea]++
	}

	if ideaCount["good"] == 0 {
		return "Fail!"
	}
	if ideaCount["good"] == 1 || ideaCount["good"] == 2 {
		return "Publish!"
	}
	return "I smell a series!"

}

func main() {
	fmt.Println(Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"})) // To(Equal("I smell a series!"))

	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"})) // .To(Equal("Publish!"))

	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad"})) // .To(Equal("Fail!"))
}
