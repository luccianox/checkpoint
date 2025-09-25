/*
findprevprime
Instructions

Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

If there are no primes inferior to the int passed as parameter the function should return 0.
Expected function
*/
package main

func FindPrevPrime(nb int) int {

	if nb < 2 {

		return 0
	}
	for i := nb; i >= 2; i-- {

		if isPrime(i) {

			return i
		}
	}
	return 0
}
func isPrime(n int) bool {

	if n <= 1 {

		return false
	}
	if n == 2 {

		return true
	}

	if n%2 == 0 {

		return false
	}
	for i := 3; i*i <= n; i += 2 {

		if n%i == 0 {
			return false
		}
	}
	return true
}
