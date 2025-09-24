/*fprime
Instructions

Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

Factors must be displayed in ascending order and separated by *.

If the number of arguments is different from 1, if the argument is invalid,
or if the integer does not have a prime factor, the program displays nothing.*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		fprime(n)
	}
}

func fprime(n int) {
	if n <= 1 {
		return
	}
	for i := 2; i < n; i++ {
		for n%i == 0 {
			fmt.Print(i)
			n /= i
		}
		if i > n {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
