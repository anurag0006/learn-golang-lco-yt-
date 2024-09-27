package main

import (
	"fmt"
	"io"
	"net/http"
)


func main()  {
	fmt.Println("Welcome")

	PerformGetReq()
}



func PerformGetReq(){
	const myurl = "http://localhost:8000"


	res, err := http.Get(myurl)

	if err !=  nil{
		panic(err)
	}


	defer res.Body.Close()

	fmt.Println("Status Code : ",res.StatusCode)
	fmt.Println("Status Code : ",res.ContentLength)


	content, _  := io.ReadAll(res.Body)

	fmt.Println(string(content))
}