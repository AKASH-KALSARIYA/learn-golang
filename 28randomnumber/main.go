package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random crypto

	myrandomnum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrandomnum)
}
