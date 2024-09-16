package main

import "fmt"

func main() {

	var x int
	fmt.Print("Enter the first number:")
	fmt.Scan(&x)
	var y int
	fmt.Print("Enter the second number:")
	fmt.Scan(&y)
	fmt.Print(x + y)

}
