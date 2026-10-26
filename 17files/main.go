package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files HE Bhai!!")

	//create a file
	contant := "This Is My File Text Bhai!!"

	file, err := os.Create("./MyFirstFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checknillerror(err)

	length, err := io.WriteString(file, contant)
	checknillerror(err)
	fmt.Println("Length of File is: ", length)

	readFile("./MyFirstFile.txt")

	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checknillerror(err)
	fmt.Println("Text data inside a file is:\n", string(databyte))
}

func checknillerror(err error) {
	if err != nil {
		panic(err)
	}
}
