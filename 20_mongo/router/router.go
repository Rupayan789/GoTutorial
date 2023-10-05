package router

import (
	"netflix_api_mock/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovieController).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controller.MarkMovieWatchedController).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controller.DeleteOneMovieController).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteAllMovieController).Methods("DELETE")

	return r
}
