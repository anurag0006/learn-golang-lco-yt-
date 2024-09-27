package main

import "fmt"

func main(){
	fmt.Println("Hello from loops")


	days := []string{"Sunday","Monday","Tuesday","Wednesday"}

	fmt.Println(days)


	for d := 0; d < len(days); d++{
            fmt.Println(days[d])
	}


	for index := range(days){
		fmt.Println(days[index])
	}



	for index,day :=range days{
		fmt.Println(index,day)
	}

	


	 
}