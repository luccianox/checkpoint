/*addprimesum
Instructions

Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.*/
package main

import(

	"os"
	"strconv"
	"fmt"
)
func isPrime(n int)bool{

}
func ()main{
	//check number of arguments
	if len(os.Args) != 2{
		fmt.Println(0)
		return
	}
	//convert os.Args from string to int
	n, err := strconv.Atoi(os.Args[1])
	//remove negatives
	if err != nil || n <= 0{
		fmt.Println(0)
		return
	}
	sum := 0
	for i := 2; i <= n{

		if isPrime(i){

			sum += 1
		}
	}
	fmt.Println(sum)
}
