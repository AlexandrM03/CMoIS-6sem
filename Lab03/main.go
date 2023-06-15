package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPrimes(100))

	fmt.Printf("gcd2(399, 433) = %d\n", gcd2(399, 433))
	fmt.Printf("gcd3(150, 430, 115) = %d\n", gcd3(150, 430, 115))
}
