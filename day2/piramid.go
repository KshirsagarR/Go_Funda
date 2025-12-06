package main

import "fmt"

func main() {
	// 	add()
	// 	sub()
	// 	multi()
	// 	div()
	n := 4
	for i := 1; i <= n; i++ {

		for j := n - 1; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
