package main

import "fmt"

func main() {
	fmt.Println("Hello From Structs!!")
	//No inheritance in golang ; No super or parent
	akash := User{"Akash", "akashkalsariya88@gmail.com", true, 16}
	fmt.Println(akash)
	fmt.Printf("akash details is:%+v", akash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
