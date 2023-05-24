package routes

import (
	"github.com/davoodmossgreen/books-i/internal/controllers"
	"github.com/gorilla/mux"
)



var RegisterBookRoutes = func(router *mux.Router){
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/add", controllers.CreateBook).Methods("POST", "GET")
	router.HandleFunc("/mybooks", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/deleteBook", controllers.DeleteBook).Methods("POST", "GET")
	router.HandleFunc("/notes", controllers.Notes).Methods("POST", "GET")
	router.HandleFunc("/deleteNote", controllers.DeleteNote).Methods("POST")
	router.HandleFunc("/history", controllers.History).Methods("GET")
}