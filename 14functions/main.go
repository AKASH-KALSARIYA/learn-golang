package main

import "fmt"

func main() {
	fmt.Println("This Is Functions!!")

	greeter()
	greetersecound()

	result := adder(3, 5)
	fmt.Println("result is: ", result)

	prores, mymessage := proadder(2, 5, 8)
	fmt.Println("Pro Result is: ", prores)
	fmt.Println("Pro message is: ", mymessage)
}

func adder(i int, j int) int {
	return i + j
}

func proadder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello Form Golang Bhau!!!"
}

func greetersecound() {
	fmt.Println("Namste Bhau!!")
}

func greeter() {
	fmt.Println("Hello From Golang")
}
