package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("SLices ssssss!!")

	var fruitList = []string{"Potato", "Tomato", "Peach"}
	// fmt.Printf("Type Of FruitList is:%T", fruitList)

	fruitList = append(fruitList, "Banana", "Mango")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 123
	highscore[2] = 456
	highscore[3] = 567
	//highscore[4] = 5654

	highscore = append(highscore, 555, 666, 777)

	// fmt.Println(highscore)

	sort.Ints(highscore)
	// fmt.Println(highscore)
	// fmt.Println(sort.IntsAreSorted(highscore))

	//how to remove a value from slice based on index
	var courses = []string{"ReactJs", "Javascript", "Python", "Ruby "}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
