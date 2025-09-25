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
