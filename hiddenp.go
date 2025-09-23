/*hiddenp
Instructions
Write a program named hiddenp that takes two strings as arguments. The program should check if the first string s1
 is hidden in the second s2. s1 is considered hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1, but not necessarily consecutively.

If s1 is hidden in s2, the program should display 1 followed by a newline.
If s1 is not hidden in s2, the program should display 0 followed by a newline.
If s1 is an empty string, it is considered hidden in any string.
If the number of arguments is different from 2, the program should display nothing.*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 3 {

		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])

	if len(s1) == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')

	}
	i := 0
	j := 0

	for i < len(s1) && j < len(s2) {

		if s1[i] == s2[2] {

			i++
		}
		j++
	}
	if i == len(s1) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
