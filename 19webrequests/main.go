package main

import (
	"fmt"
	"io"
	"net/http"
)


func main()  {
	resp, err := http.Get("https://www.funworldblr.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.StatusCode)

	databytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))


	
	defer resp.Body.Close()


}