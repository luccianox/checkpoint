/*
repeatalpha
Instructions

Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetica
l character as many times as its alphabetical index.

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
Expected Function
*/
package main

import "fmt"

func RepeatAlpha(s string) string {

	//loop through s
	result := ""
	for i := 0; i < len(s); i++ {
		//look for small letters
		char := s[i]
		if s[i] >= 'a' && s[i] <= 'z' {

			index := char - 'a' + 1
			for j := 0; j < int((index)); j++ {

				result += string(char)
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {

			index := char - 'A' + 1
			for j := 0; j < int((index)); j++ {

				result += string(char)
			}
		} else {
			result += string(char)
		}

	}
	return result
}
func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
