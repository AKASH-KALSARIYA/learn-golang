package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("WelCome to our Pizza App")
	fmt.Println("Please Rate a Pizza 1 to 5")

	Reader := bufio.NewReader(os.Stdin)

	input, _ := Reader.ReadString('\n')

	fmt.Println("Thanks For Rating : ", input)
	//conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to Your rating: ", numRating+1)
	}

}
