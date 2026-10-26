package main

import "fmt"

func main() {
	// loginCount := 23
	var result string

	// if loginCount < 10 {
	// 	result = "Regular user"
	// } else {
	// 	result = "somthing else"
	// }

	// fmt.Println(result)

	if num := 22; num > 10 {
		result = "greater than 10"
	} else {
		result = "less than 10"
	}
	fmt.Println(result)
}
