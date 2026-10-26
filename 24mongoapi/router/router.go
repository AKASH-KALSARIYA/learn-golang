package Router

import (
	"github.com/AKASH-KALSARIYA/mygolang/24mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllmovies).Methods("Get")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("Post")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("Put")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("Delete")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("Delete")

	return router
}
