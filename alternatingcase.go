/*
altERnaTIng cAsE <=> ALTerNAtiNG CaSe
altERnaTIng cAsE <=> ALTerNAtiNG CaSe
Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:

ToAlternatingCase("hello world"); // => "HELLO WORLD"
ToAlternatingCase("HELLO WORLD"); // => "hello world"
ToAlternatingCase("hello WORLD"); // => "HELLO world"
ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
As usual, your function/method should be pure, i.e. it should not mutate the original string.
*/

package main

import (
	"fmt"
	"strings"
)

func ToAlternatingCase(str string) string {
	var result string
	for _, rune := range str {
		value := string(rune)
		if isLetter(value) {
			if isLowerCase(value) {
				result += strings.ToUpper(value)
			}

			if !isLowerCase(value) {
				result += strings.ToLower(value)
			}
		} else {
			result += value
		}
	}

	return result
}

func isLetter(input string) bool {
	r := []rune(strings.ToLower(input))
	return r[0] >= 97 && r[0] <= 122
}

func isLowerCase(input string) bool {
	r := []rune(input)
	return r[0] >= 97 && r[0] <= 122
}

func main() {
	fmt.Println(ToAlternatingCase("hello world"))                        // => "HELLO WORLD"
	fmt.Println(ToAlternatingCase("HELLO WORLD"))                        // => "hello world"
	fmt.Println(ToAlternatingCase("hello WORLD"))                        // => "HELLO world"
	fmt.Println(ToAlternatingCase("HeLLo WoRLD"))                        // => "hEllO wOrld"
	fmt.Println(ToAlternatingCase("12345"))                              // => "12345" (Non-alphabetical characters are unaffected)
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))                         // => "1A2B3C4D5E"
	fmt.Println(ToAlternatingCase("String.prototype.toAlternatingCase")) // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
}
