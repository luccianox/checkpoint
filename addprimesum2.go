package main

import (
	"fmt"
	"os"
)

func toInt(s string) (int, bool) {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false // not a digit
		}
		n = n*10 + int(ch-'0')
	}
	return n, true
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	n, ok := toInt(os.Args[1])
	if !ok || n <= 0 {
		fmt.Println(0)
		return
	}

	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
