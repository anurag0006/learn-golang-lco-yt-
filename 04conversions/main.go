package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main()  {
	fmt.Println("Rate  the pizza:")
	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else{
		numRating = numRating+1
	}

	fmt.Println(numRating)

}