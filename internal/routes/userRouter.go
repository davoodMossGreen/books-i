package routes

import (
	"github.com/davoodmossgreen/books-i/internal/controllers"
	"github.com/gorilla/mux"
)



var RegisterUserRoutes = func(router *mux.Router){
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST", "GET")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST", "GET")
	router.HandleFunc("/logout", controllers.LogOut).Methods("GET")
	router.HandleFunc("/search", controllers.GetUserByName).Methods("GET", "POST")
	router.HandleFunc("/deleteAccount", controllers.DeleteUser).Methods("POST", "GET")
}
