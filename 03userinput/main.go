package main

import (
	"bufio"
	"fmt"
	"os"
)


func main()  {
	welcome := "Hello user!, welcome"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")

	input, err := reader.ReadString('\n')

	if err != nil{
       fmt.Println("Got some error:",err)
	}

	fmt.Println("Thanks for reading",input)
}