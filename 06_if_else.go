// pretty basic as far as if/else goes

package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// 7 is odd

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// 8 is divisible by 4

	// will only run the first successful condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// 9 has 1 digit
}

// note: there is no ternary if in Go ( ?: )
