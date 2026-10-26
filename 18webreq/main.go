package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://hirax.in/"

func main() {
	fmt.Println("Hello Web Req!!")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Of This Type is:%T\n", response)

	defer response.Body.Close() //caller responcibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	context := string(databytes)
	fmt.Println(string(context))
}
