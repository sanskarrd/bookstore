package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanskardubey/Documents/Projects/gocode/src/go-bookstore/pkg/controllers"
)
 
//RegisterBookStoreRoutes is a function that registers the routes for the bookstore
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}