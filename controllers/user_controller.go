package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChanchalS7/golang-rbac/auth"
	"github.com/ChanchalS7/golang-rbac/models"
	"github.com/ChanchalS7/golang-rbac/services"
	"github.com/ChanchalS7/golang-rbac/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
//GetUsers handles fetching paginated users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10 //Default limit
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0 //Default offset
	}
	users, err := uc.UserService.GetAllUsers(limit, offset)
	if err != nil {
		utils.ResponseWithError(w, http.
			StatusInternalServerError, err.Error())
			return
	}
	utils.RespondWithJSON(w, http.StatusOK,users)
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


// SignUp handles user registrations

func (uc *UserController) SignUp(w http.ResponseWriter, r *http.Request){
	var user models.User

	_= json.NewDecoder(r.Body).Decode(&user)

	//Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error hashing password")
		return 
	}
	user.Password = string(hashedPassword)

	_, err  = uc.UserService.CreateUser(user)

	if err !=nil {
		utils.ResponseWithError(w, http.StatusInternalServerError,"Error creating user")
		return 
	}
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

//Login handles user authentication 
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {

	var creds models.User

	_ = json.NewDecoder(r.Body).Decode(&creds)

	//Fetch user from database

	user, err := uc.UserService.FindUserByEmail(creds.Email)

	if err != nil {
		if err == mongo.ErrNoDocuments{
			utils.ResponseWithError(w, http.StatusUnauthorized, "Invalid credentials")
		}else {
			utils.ResponseWithError(w, http.StatusInternalServerError, "Error finding user")
		}
		return 
	}

	//Compare passwords

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		utils.ResponseWithError(w, http.StatusUnauthorized,"Invalid credentials")
		return 
	}

	//Generate JWT
	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token":token})
}