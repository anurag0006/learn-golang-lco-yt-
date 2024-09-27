package main

import "fmt"

func main()  {
	fmt.Println("hello from pointers!")

	var one int = 2
	var ptr  *int = &one

	fmt.Println(ptr)


}