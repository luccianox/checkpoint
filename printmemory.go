<<<<<<< HEAD
/*Instructions
Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters.
The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed,
or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
Expected function*/

=======
/*
printmemory
Instructions

Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans,
present on the ASCII encoding.
Expected function
*/
>>>>>>> ffbb1941cbe80c2c3a716c29c8ce1bf9c71b5be0
package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
<<<<<<< HEAD

	//loop through arr
	for i := 0; i <= arr; i++ {
		high := arr[i] / 16
		low := arr[i] % 16
		//print high
=======
	// loop through arr

	for i := 0; i < len(arr); i++ {

		high := arr[i] / 16

		low := arr[i] % 16

		// print high
>>>>>>> ffbb1941cbe80c2c3a716c29c8ce1bf9c71b5be0

		if high < 10 {
			z01.PrintRune(rune('0' + high))
		} else {
			z01.PrintRune(rune('a' + (high - 10)))
		}
		if low < 10 {
			z01.PrintRune(rune('0' + low))
		} else {
			z01.PrintRune(rune('a' + (low - 10)))
		}
		if (i+1)%4 == 0 {
<<<<<<< HEAD

			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
	//Print ASCII characters

	for _, ch := range arr {

		if ch >= 32 && ch <= 126 {
			z01.PrintRune(rune(ch))
=======
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	// print ASCII
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
>>>>>>> ffbb1941cbe80c2c3a716c29c8ce1bf9c71b5be0
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
<<<<<<< HEAD
=======

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
>>>>>>> ffbb1941cbe80c2c3a716c29c8ce1bf9c71b5be0
