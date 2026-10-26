package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition!!")

	waitgroup := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score []int

	waitgroup.Add(4)
	go func(waitgroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		waitgroup.Done()
	}(waitgroup, mut)
	// waitgroup.Add(1)
	go func(waitgroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		waitgroup.Done()
	}(waitgroup, mut)
	go func(waitgroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		waitgroup.Done()
	}(waitgroup, mut)
	go func(waitgroup *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Four R")
		mut.Lock()
		fmt.Println(score)
		mut.Unlock()
		waitgroup.Done()
	}(waitgroup, mut)

	waitgroup.Wait()

	fmt.Println(score)
}
