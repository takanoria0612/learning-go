package main

import "fmt"

func main() {
	var x = 100
	fmt.Println("x address:\t", &x)

	var y *int
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)
	y = &x
	fmt.Println("y value:\t", y)
	fmt.Println("y address:\t", &y)
}
