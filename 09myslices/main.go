package main

import "fmt"

func main() {
	fmt.Println("Hello from slices!")

	var fruitList = []string{"Apple"}

	// Use fmt.Printf for formatted output
	fmt.Printf("%T\n", fruitList)

	fruitList = append(fruitList, "Mango","Banana")


	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)


	

}
