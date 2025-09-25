/*Instructions
Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters.
The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed,
or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
Expected function*/

/*
printmemory
Instructions

Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans,
present on the ASCII encoding.
Expected function
*/

package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {

	for i := 0; i < len(arr); i++ {
		high := arr[i] / 16
		low := arr[i] % 16
		//print high
		if high < 10 {

			z01.PrintRune(rune('0' + high))
		} else {
			z01.PrintRune(rune('a' + (high - 10)))
		}
		//print low

		if low < 10 {
			z01.PrintRune(rune('0' + low))
		} else {
			z01.PrintRune(rune('a' + (low - 10)))
		}
		z01.PrintRune(' ')
		if (i+1)%4 == 0 {

			z01.PrintRune('\n')
		}
	}
	if len(arr)%4 != 0 {
		z01.PrintRune('\n')
	}
	//Print ASCII
	for _, ch := range arr {

		if ch >= 32 && ch <= 126 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
