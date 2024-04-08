package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/idugan100/fitness_server_330/database"
	"github.com/idugan100/fitness_server_330/models"
)

type AuthController struct {
	UserRepo database.UserRepository
}

func NewAuthController(r database.UserRepository) AuthController {
	return AuthController{UserRepo: r}
}

func (a AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyString, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyString, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = a.UserRepo.Signup(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	//read crednetials from request
	bodyString, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//turn it into a struct
	var user models.User
	err = json.Unmarshal(bodyString, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//try to log the user in and return the full user if so
	user, err = a.UserRepo.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, _ := json.Marshal(user)
	w.Write(data)
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	//end session
	w.Write([]byte("logout"))
}

func (a AuthController) AllUsers(w http.ResponseWriter, r *http.Request) {
	//must be admin
	userList, err := a.UserRepo.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(userList)
	w.Write(data)
}

func SetUserMiddleware(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	// sets the user in the request context
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("setting user context")
		//if user is not logged in return 403 else attatch user to requst context
		fn(w, r)
	}
}
