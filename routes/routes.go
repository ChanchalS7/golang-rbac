package routes

import (
	"github.com/gorilla/mux"
	"github.com/ChanchalS7/golang-rbac/controllers"

)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(userController *controllers.UserController) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/users", userController.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
    router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

    return router
}