package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps Bhau!!!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "RUBY"
	languages["PY"] = "Python"

	fmt.Println("List of all languges:", languages)
	fmt.Println("JS is Short for: ", languages["JS"])
	fmt.Println("RB is Short for: ", languages["RB"])
	fmt.Println("PY is Short for: ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops in Golang
	for key, value := range languages {
		fmt.Printf("For key %v,valuse is %v\n", key, value)
	}

}
