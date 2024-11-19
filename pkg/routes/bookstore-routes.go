package routes

import (
	"github.com/CharlieDeepk/go_basic_project/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{booId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{booId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{booId}", controllers.DeleteBook).Methods("DELETE")

}
