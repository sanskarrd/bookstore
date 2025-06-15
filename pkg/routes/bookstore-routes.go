package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanskardubey/Documents/Projects/gocode/src/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.Handlefunc("/book/", controllers.CreateBook).Methods("POST")
	router.Handlefunc("/book", controllers.GetBook).Methods("GET")
	router.Handlefunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.Handlefunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.Handlefunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}