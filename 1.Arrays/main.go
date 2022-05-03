package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	myarray := []int{1, 2, 3, 4, 5, 56, 6, 7, 7}

	fmt.Println(myarray[7])
	for i, n := range myarray {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(n)
	}
}
