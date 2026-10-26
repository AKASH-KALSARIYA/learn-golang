package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Akash is welcoming you!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Rating For Our Shop:")
	//comma ok or err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks For Rating,", input)
	fmt.Printf("type of that rating is %T", input)
}
