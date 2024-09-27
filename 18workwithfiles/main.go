package main

import (
	"fmt"
	"io"
	"os"
)


func main()  {
	fmt.Println("Welcome to files in golang!")
	content := "Need to be added in a text file"

	file,err := os.Create("./myfile.txt")
	checkNilErr(err)
	// if err != nil{
	// 	panic(err)
	// }

	length ,err := io.WriteString(file,content)

	checkNilErr(err)
	fmt.Println(length)

	defer file.Close()
	ReadFile("./myfile.txt")
	
}



func ReadFile(filename string){
   databyte, err := os.ReadFile(filename)
   checkNilErr(err)
  fmt.Println(string(databyte))
}


func  checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}