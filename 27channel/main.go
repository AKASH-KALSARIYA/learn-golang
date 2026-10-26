package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels!!")

	myCh := make(chan int, 3)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(3)
	//that go routins Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, IsChannelOpen := <-myCh
		fmt.Println(IsChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	wg.Done()
	// }(myCh, wg)

	//that go-routins is send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
