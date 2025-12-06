package main

import "fmt"

//import "math"

// func main() {
// 	add()
// 	sub()
// 	multi()
// 	div()
// }

var a int = 4
var b int = 6
var r1 int

func add() {
	r1 = a + b
	fmt.Println("Addition: ", r1)

}
func sub() {
	r1 = a - b
	fmt.Println("Substraction: ", r1)
}
func multi() {
	r1 = a * b
	fmt.Println("Multiplication: ", r1)
}
func div() {
	r1 = a % b
	fmt.Println("Division: ", r1)

}
