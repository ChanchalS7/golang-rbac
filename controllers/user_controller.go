package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ChanchalS7/golang-rbac/models"
	"github.com/ChanchalS7/golang-rbac/services"
	"github.com/ChanchalS7/golang-rbac/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserController struct{
	UserService *services.UserService
}

//CreateUser handles the creation of a new user
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_= json.NewDecoder(r.Body).Decode(&user)

	user.ID = primitive.NewObjectID()

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}
//GetUserByID handles fetching a user by their ID
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	user, err:= uc.UserService.GetUserByID(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, err.Error())
		return 
	}
	utils.RespondWithJSON(w, http.StatusOK, user)


}

//UpdateUser handles updating a user's details
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	updatedUser, err := uc.UserService.UpdateUser(params["id"], user)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedUser)
}

//DeleteUser handles deleting a user by thier ID
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	err := uc.UserService.DeleteUser(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return 
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result":"success"})
}