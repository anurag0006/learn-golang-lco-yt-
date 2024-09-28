package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anurag0006/apiwithdb/router"
)


func main()  {
	fmt.Println("Server is Starting...")
    r := router.Router()
	fmt.Println("Listening on port 4000...")

	log.Fatal(http.ListenAndServe(":4000",r))


}