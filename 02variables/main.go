package main

import "fmt"

const logintoken string = "akashtoken" //public

func main() {
	var username string = "Akash"
	fmt.Println(username)
	fmt.Printf("variables is of type:%T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variables is of type:%T \n", isLoggedin)

	var smallval int32 = 355
	fmt.Println(smallval)
	fmt.Printf("variables is of type:%T \n", smallval)

	var smallfloat float64 = 355.232323232323
	fmt.Println(smallfloat)
	fmt.Printf("variables is of type:%T \n", smallfloat)

	//default values and eclaration
	var notgarbage int
	fmt.Println(notgarbage)
	fmt.Printf("the type of var :%T\n", notgarbage)

	//implcit way/not type declaration
	var web = "akash.in"
	println(web)

	//no var style
	numberofuser := 3555
	fmt.Println(numberofuser)

	fmt.Println(logintoken)
	fmt.Printf("the type of var :%T\n", logintoken)
}
