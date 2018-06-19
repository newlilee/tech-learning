// algorithm_gcd.go
package main

import (
	"fmt"
)

func gcd(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}

func main() {
	retVal := gcd(16, 4)
	fmt.Println("The gcd is:", retVal)
}