package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string
	Price int
	Platform string
	Tags []string
}



func main()  {
	jsonData  := EncodeJson()

	decodeJson(jsonData)





}


func EncodeJson()string{
	courses := []course{
       {"Reactjs",299,"udemy",[]string{"web-dev","js"}},
       {"mern",199,"youtube",[]string{"mern","js","react"}},
       {"angular",1000,"twitch",nil},
	}


	// package this data as json:

	finalJson,err := json.MarshalIndent(courses,"","\t")

	if err != nil{
		panic(err)
	}

	return string(finalJson)
}


func decodeJson(jsonData string){
	jsonDataFromWeb := jsonData

	fmt.Println("hello")

	fmt.Println(jsonDataFromWeb)


	fmt.Println(json.Valid([]byte(jsonDataFromWeb)))
    
	var newcourses course

	json.Unmarshal([]byte(jsonDataFromWeb),&newcourses)

	fmt.Println("%#v\n",new)
}