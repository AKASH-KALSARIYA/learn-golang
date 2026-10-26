package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello This Is My Modules!!!(MOD)")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("Get")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("HEllo There are Mod Users!!!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
