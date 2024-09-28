package router

import (
	controller "github.com/anurag0006/apiwithdb/controllers"
	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	 router := mux.NewRouter()
     
	 router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	 router.HandleFunc("/api/movie",controller.CreateNewMovie).Methods("POST")
	 router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	 router.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	 router.HandleFunc("/api/deleteall",controller.DeleteAllMovies).Methods("DELETE")


	 return router
}