package main

import "fmt"

func main() {
	defer fmt.Println("Hello This Is A Defer KeyWord!!")
	defer fmt.Println("first")
	defer fmt.Println("secound")
	defer fmt.Println("Third")
	fmt.Println("Hello This Is A Defer!!")
	MyDefer()
}

func MyDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
