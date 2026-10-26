package main

import "fmt"

func main() {
	fmt.Println("Pointer Concepts")

	myNumber := 22

	var ptr = &myNumber

	fmt.Println("value of actual pointer is:", ptr)
	fmt.Println("value of actual pointer is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value Is: ", myNumber)
}
