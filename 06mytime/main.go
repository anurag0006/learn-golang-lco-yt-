package main

import (
	"fmt"
	"time"
)


func main()  {
	fmt.Println("Time study")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-06 15:04:05 Monday"))


	createdDate := time.Date(2020, time.August,12,11,11,11,0,time.Local)  

	fmt.Println(createdDate.Format("01-02-06"))
}