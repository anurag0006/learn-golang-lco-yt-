package main

import "fmt"

func main()  {
	defer fmt.Println("WOrld 1!")
	defer fmt.Println("WOrld 2!")
	defer fmt.Println("WOrld 3!")
	fmt.Println("WOrld!")
}