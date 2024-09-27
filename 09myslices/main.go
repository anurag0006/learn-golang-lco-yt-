package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from slices!")

	var fruitList = []string{"Apple"}

	// Use fmt.Printf for formatted output
	fmt.Printf("%T\n", fruitList)

	fruitList = append(fruitList, "Mango","Banana")


	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)


	// Removing from slicee:

	var courses = make([]string, 3)	
	courses[0] = "ma101"
	courses[1] = "js"
	courses[2] = "react"
	fmt.Println(courses)

	var index int = 1

	courses = append(courses[:index],courses[index+1:]...)
    fmt.Println(courses)




}
