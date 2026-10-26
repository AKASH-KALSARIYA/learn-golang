package main

import (
	"fmt"
)

func main() {
	fmt.Println("These were loops!!!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thurshday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); ++i {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rouguevalue := 1

	for rouguevalue < 10 {

		if rouguevalue == 2 {
			goto lco
		}

		if rouguevalue == 5 {
			rouguevalue++
			continue
		}

		fmt.Println("Value is: ", rouguevalue)
		rouguevalue++
	}

	//GoTo statement
lco:
	fmt.Println("Jumping at Google.com")
}
