package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET REQUEST KIYA HE!!!")
	// getreq()
	// postjsonreq()
	postFormreq()
}

func getreq() {
	const myurl = "http://localhost:8000/get"
	responce, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	fmt.Println("Status Code: ", responce.StatusCode)
	fmt.Println("Content Length: ", responce.ContentLength)

	//we can print all of things using Builder method or simple method whichever we used in comments in the scenario that prints same things
	//first way
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	//secound way
	// content, _ := ioutil.ReadAll(responce.Body)
	// fmt.Println(string(content))
}

func postjsonreq() {
	const myurl = "http://localhost:8000/post"

	//fake Json PayLoad

	requestBody := strings.NewReader(`
		{
			"CourseName" : "Golang Karna He!!",
			"Price" : 846,
			"PlateForm" : "Google.com" 
		}
	`)
	responce, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()
	content, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(content))
}

func postFormreq() {
	const myurl = "http://localhost:8000/postform"

	//FormData
	data := url.Values{}
	data.Add("Firstname", "Akash")
	data.Add("Lastname", "Kalsariya")
	data.Add("Email", "test@gmail.com")

	responce, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	content, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(content))
}
