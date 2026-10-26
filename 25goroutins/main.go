package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal = []string{"test"}

var waitgroup sync.WaitGroup //pointers

var mutex sync.Mutex //pointer

func main() {
	fmt.Println("GO ROUTINs!")
	// go Greeter("HELLO")
	// Greeter("MY FRIEND!!")
	website := []string{
		"http://google.com",
		"http://github.com",
	}
	for _, web := range website {
		go getstatuscode(web)
		waitgroup.Add(1)
	}
	waitgroup.Wait()
	fmt.Println(signal)

}

// func Greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getstatuscode(endpoint string) {
	defer waitgroup.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOps endpoint")
	}
	mutex.Lock()
	signal = append(signal, endpoint)
	mutex.Unlock()
	fmt.Printf("%d 200 statuscodde for website %s\n ", res.StatusCode, endpoint)
}
