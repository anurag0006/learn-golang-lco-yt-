package main

import "fmt"


func main()  {
	fmt.Println("Structs!")


	// no inheritance in golang, neither is super or parent;

	anurag := User{"Anurag","A@gmail.com",true,22}

	fmt.Println(anurag)


}


type User struct{
	Name string
	Email string
	Status bool
	Age int
}