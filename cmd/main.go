package main

import (
	"log"
	"net/http"

	"github.com/ChanchalS7/golang-rbac/configs"
	"github.com/ChanchalS7/golang-rbac/controllers"
	"github.com/ChanchalS7/golang-rbac/repositories"
	"github.com/ChanchalS7/golang-rbac/routes"
	"github.com/ChanchalS7/golang-rbac/services"
)

func main(){
	configs.Loadenv()
	client := configs.ConnectDB()

		userRepo := &repositories.UserRepository{Collection: client.Database("golang-rbac").Collection("users")}
		userService := &services.UserService{Repo : userRepo}
		userController := &controllers.UserController{UserService: userService}
		
		router := routes.InitializeRoutes(userController)
		log.Println("Server started at: 8080")
		log.Fatal(http.ListenAndServe(":8080",router))
}