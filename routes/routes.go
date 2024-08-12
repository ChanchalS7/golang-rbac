package routes

import (
	"github.com/ChanchalS7/golang-rbac/auth"
	"github.com/ChanchalS7/golang-rbac/controllers"
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(userController *controllers.UserController) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/signup",userController.SignUp).Methods("POST")
    router.HandleFunc("/login", userController.Login).Methods("POST")

    //Protected routes
    api:= router.PathPrefix("/api").Subrouter()
    api.Use(auth.AuthMiddleware)
    router.HandleFunc("/users", userController.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
    router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

    return router
}