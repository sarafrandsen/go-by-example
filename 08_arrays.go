// arrays in Go are fixed length!
package main

import "fmt"

func main() {

	var a [5]int // this arr will hold 5 ints
	// arrs by default are zero-valued (for ints, this means 0)
	fmt.Println("emp:", a) // emp: [0 0 0 0 0]

	// set the value at index 4 to 100
	a[4] = 100
	// new value of 'a' array
	fmt.Println("set", a) // set [0 0 0 0 100]
	// new value at index 4
	fmt.Println("get:", a[4]) // get: 100

	// length
	fmt.Println("len:", len(a)) // len: 5

	// declare and initialize array in single line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) // dcl: [1 2 3 4 5]

	// arrays are technically one-dimensional, but can be composed to be multi-dimensional (need clarification on wording????)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0 1 2] [1 2 3]]

}
