package main

import "fmt"

func main() {
	fmt.Println("Hello From Structs!!")
	//No inheritance in golang ; No super or parent
	akash := User{"Akash", "akashkalsariya88@gmail.com", true, 16}
	fmt.Println(akash)
	fmt.Printf("akash details is:%+v\n", akash)
	akash.GetStatus()
	akash.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("IS User Is Acive:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
