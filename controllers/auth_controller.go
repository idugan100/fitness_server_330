package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/idugan100/fitness_server_330/database"
	"github.com/idugan100/fitness_server_330/models"
)

var mySigningKey = []byte("AllYourBase")

type JWTBody struct {
	UserId  int  `json:"userID"`
	IsAdmin bool `json:"isAdmin"`
	jwt.RegisteredClaims
}
type AuthController struct {
	UserRepo database.UserRepository
}

func NewAuthController(r database.UserRepository) AuthController {
	return AuthController{UserRepo: r}
}

func (a AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyString, err := io.ReadAll(r.Body)
	defer r.Body.Close()

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
	token, err := CreateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(token))
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	//read crednetials from request
	bodyString, err := io.ReadAll(r.Body)
	defer r.Body.Close()

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

	token, err := CreateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
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

		auth_header := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth_header, "Bearer ") {
			http.Error(w, "invalid authorization header format", http.StatusBadRequest)
			return
		}

		user, err := ParseToken(strings.TrimPrefix(auth_header, "Bearer "))
		if !strings.HasPrefix(auth_header, "Bearer ") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		new_ctx := context.WithValue(ctx, "user", user)
		new_r := r.WithContext(new_ctx)
		fn(w, new_r)
	}
}

func RequireAdmin(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		fn(w, r)
	}
}

func CheckPermission(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		id, _ := strconv.Atoi(r.PathValue("userID"))
		if user.IsAdmin {
			fn(w, r)

		} else if user.Id != id {
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			fn(w, r)
		}
	}
}

func CreateToken(user models.User) (string, error) {
	claims := JWTBody{
		user.Id,
		user.IsAdmin,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseToken(token string) (models.User, error) {
	var user models.User
	t, err := jwt.ParseWithClaims(token, &JWTBody{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return user, err
	}
	user.IsAdmin = t.Claims.(*JWTBody).IsAdmin
	user.Id = t.Claims.(*JWTBody).UserId

	return user, nil
}
