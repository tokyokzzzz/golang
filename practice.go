package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Println("write a number")
	fmt.Scan(&a)
	isPrime := true
	for i := 2; i < a; i++ {
		if a%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(a, "is a prime number")
	} else {
		fmt.Println(a, "is not a prime number")
	}
}
