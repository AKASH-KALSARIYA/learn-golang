package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=reactjs&rlz=1C1GCEU_en&oq=reactjs&gs_lcrp=EgZjaHJvbWUqCggAEAAYsQMYgAQyCggAEAAYsQMYgAQyDQgBEAAYgwEYsQMYgAQyBwgCEAAYgAQyBwgDEAAYgAQyBwgEEAAYgAQyBwgFEAAYgAQyBggGEEUYPDIGCAcQRRg80gEIMjM0NGowajeoAgCwAgA&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("This Is url Handling!!!")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	QueryParams := result.Query()
	fmt.Printf("The Type Of query Params is: %T\n", QueryParams)

	fmt.Println(QueryParams["coursename"])

	for _, val := range QueryParams {
		fmt.Println("Param is: ", val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "Ico. dev",
		Path:    "/tutess",
		RawPath: "user=hitesh",
	}
	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}
