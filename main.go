package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter ur name")
	fmt.Scanln(&name)

	fmt.Print("enter ur age")
	fmt.Scanln(&age)

	fmt.Print("Hello", name, "ur age is", age)
}
