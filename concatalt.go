/*
concatalternate
Instructions
Write a function ConcatAlternate() that receives two slices of an int as arguments and returns

	a new slice with the result of the alternated values of each slice.

The input slices can be of different lengths.
The new slice should start with an element of the largest slice.
If the slices are of equal length, the new slice should return the elements of the first slice
first and then the elements of the second slice.
Expected function
*/
package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {

	result := []int{}
	if len(slice2) > len(slice1) {

		slice1, slice2 = slice2, slice1
	}

	for i, v := range slice1 {
		result = append(result, v)
		if i < len(slice2) {
			result = append(result, slice2[i])
		}
	}
	return result
}
func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
