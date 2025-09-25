/*
union
Instructions
Write a program that takes two string and displays, without doubles, the characters that appear in either one

	of the string.The display will be in the same order that the characters appear on the command line
	 and will be followed by a newline ('\n').

If the number of arguments is different from 2, then the program displays a newline ('\n').
*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	union := s1 + s2
	result := []rune{}

	// loop through union
	for _, ch := range union {

		found := false
		// loop through result
		for _, r := range result {

			if ch == r {

				found = true
				break
			}
		}
		if found == false {
			result = append(result, ch)
		}
	}
	for _, m := range result {

		z01.PrintRune(m)
	}
	z01.PrintRune('\n')
}
