package controllers

import (
	"fmt"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// get username, password  and create user and start session
	w.Write([]byte("signed in"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	//check if username and password are good
	w.Write([]byte("login"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//end session
	w.Write([]byte("logout"))
}

func SetUserMiddleware(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	// sets the user in the request context
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("setting user context")
		//if user is not logged in return 403 else attatch user to requst context
		fn(w, r)
	}
}
