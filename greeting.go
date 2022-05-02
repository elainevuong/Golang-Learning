/*
Make a function that will return a greeting statement that uses an input; your program should return,
"Hello, <name> how are you doing today?".
*/

package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)

}

func main() {
	fmt.Println(Greet("Elaine"))
	fmt.Println(Greet("Jesse"))
}
