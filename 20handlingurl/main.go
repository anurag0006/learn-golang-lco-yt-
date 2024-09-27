package main

import (
	"fmt"
	"net/url"
)


const myurl string = "https://www.funworldblr.com/tickets?name=anurag&age=12"

func main()  {
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)


	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["name"])
}