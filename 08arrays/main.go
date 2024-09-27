package main

import "fmt"



func main()  {
	fmt.Println("Hello from arrays!")


	var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[2] = "kamboj"
	fruitList[3] = "anurag"

	fmt.Println(len(fruitList))
	fmt.Println((fruitList))

}