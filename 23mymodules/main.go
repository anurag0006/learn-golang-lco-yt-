package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main()  {
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")


    fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":4000",r))
}



func serveHome(w http.ResponseWriter, r * http.Request){
	w.Write([]byte("<h1> Welcome to golang!</h1>"))
}