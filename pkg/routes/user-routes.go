package routes

import (
	"github.com/gorilla/mux"
	"github.com/zhbhmd/go-todo/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/user/{userId}/todo", controllers.GetUserTodo).Methods("GET")
}
