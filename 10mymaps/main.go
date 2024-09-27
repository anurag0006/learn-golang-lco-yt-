package main

import "fmt"


func main()  {
	lang := make(map[string]string)


	lang["js"] = "javascript"
	lang["rb"] = "ruby"
	lang["py"] = "python"

	fmt.Println(lang["js"])
	fmt.Println(lang)

	delete(lang,"rb")

	fmt.Println(lang)


	// loops on maps:


	for key,value  := range lang{
		fmt.Println(key,value)
	}


}