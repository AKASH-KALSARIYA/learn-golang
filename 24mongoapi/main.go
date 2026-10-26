package main

import (
	"fmt"
	"log"
	"net/http"

	Router "github.com/AKASH-KALSARIYA/mygolang/24mongoapi/router"
)

func main() {
	fmt.Println("NETFLIX API!!!")
	r := Router.Router()
	fmt.Println("Server Is Getting Started...")
	log.Fatal(http.ListenAndServe(":2004", r))
	fmt.Println("Listening At Port 2004....")

}
