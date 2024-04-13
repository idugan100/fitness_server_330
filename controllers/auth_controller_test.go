package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/idugan100/fitness_server_330/models"
)

func TestIsAdminAllow(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	token, _ := CreateToken(models.User{Id: 1, IsAdmin: true})
	r.Header.Add("Authorization", "Bearer "+token)

	s := SetUserMiddleware(RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	s(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusOK, w.Code)
	}
}

func TestIsAdminReject(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	token, _ := CreateToken(models.User{Id: 1, IsAdmin: false})
	r.Header.Add("Authorization", "Bearer "+token)

	s := SetUserMiddleware(RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	s(w, r)

	if w.Code != http.StatusForbidden {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusForbidden, w.Code)
	}
}

func TestErrorWithoutToken(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	s := SetUserMiddleware(RequireAdmin(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	s(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected response code. Expected: %d Recieved: %d", http.StatusBadRequest, w.Code)
	}
}

func TestCheckPermissionIsAdmin(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/activites", nil)
	r.SetPathValue("userID", "1")
	w := httptest.NewRecorder()
	token, _ := CreateToken(models.User{Id: 1, IsAdmin: true})
	r.Header.Add("Authorization", "Bearer "+token)

	handler := SetUserMiddleware(CheckPermission(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code for admin on check permission. Expected %d, Recieved %d", http.StatusOK, w.Code)
	}
}

func TestCheckPermissionIsUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/activites", nil)
	r.SetPathValue("userID", "1")
	w := httptest.NewRecorder()
	token, _ := CreateToken(models.User{Id: 1, IsAdmin: false})
	r.Header.Add("Authorization", "Bearer "+token)

	handler := SetUserMiddleware(CheckPermission(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code for admin on check permission. Expected %d, Recieved %d", http.StatusOK, w.Code)
	}
}

func CheckPermissionIsNotUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/activites", nil)
	r.SetPathValue("userID", "1")
	w := httptest.NewRecorder()
	token, _ := CreateToken(models.User{Id: 2, IsAdmin: false})
	r.Header.Add("Authorization", "Bearer "+token)

	handler := SetUserMiddleware(CheckPermission(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler(w, r)

	if w.Code != http.StatusForbidden {
		t.Errorf("unexpected status code for admin on check permission. Expected %d, Recieved %d", http.StatusForbidden, w.Code)
	}
}

func TestAllUsers(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/allusers", nil)
	w := httptest.NewRecorder()

	auth_controller.AllUsers(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code for all users route. Expected %d, Recieved %d", http.StatusOK, w.Code)

	}
	results, err := io.ReadAll(w.Result().Body)

	if err != nil {
		t.Errorf("unexpected error parsing json results of all users: %s", err.Error())
	}

	if !(strings.Contains(string(results), "isaac") || strings.Contains(string(results), "sally") || strings.Contains(string(results), "tom")) {
		t.Errorf("missing one of the users in results: %s", results)
	}

}

func TestInValidSignUpJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("\"username\":\"test\",\"password\":\"123\"}"))
	w := httptest.NewRecorder()

	auth_controller.Signup(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected status code for signup route. Expected %d, Recieved %d", http.StatusCreated, w.Code)
	}

	userList, _ := user_repo.All()

	if len(userList) != 3 {
		t.Errorf("user incorrectly added after invalid signup")
	}
}

func TestInValidSignUpBody(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\"password\":\"123\"}"))
	w := httptest.NewRecorder()

	auth_controller.Signup(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected status code for signup route. Expected %d, Recieved %d", http.StatusCreated, w.Code)
	}

	userList, _ := user_repo.All()

	if len(userList) != 3 {
		t.Errorf("user incorrectly added after invalid signup")
	}
}

func TestValidSignUp(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\"username\":\"test\",\"password\":\"123\"}"))
	w := httptest.NewRecorder()

	auth_controller.Signup(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("unexpected status code for signup route. Expected %d, Recieved %d", http.StatusCreated, w.Code)
	}

	userList, _ := user_repo.All()

	if len(userList) != 4 {
		t.Errorf("no user added after signup")
	}

	token, err := io.ReadAll(w.Result().Body)

	if err != nil {
		t.Errorf("unexpected error parsing data from signup token: %s", err.Error())
	}
	if string(token) == "" {
		t.Errorf("not token returned")
	}
}

func TestLoginWithInValidCredentials(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\"username\":\"tom\",\"password\":\"12\"}"))
	w := httptest.NewRecorder()
	auth_controller.Login(w, r)
	if w.Code == http.StatusOK {
		t.Errorf("user was logged in with invalid credentials")
	}
}
func TestLoginWithInValidJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("\"username\":\"tom\",\"password\":\"12\"}"))
	w := httptest.NewRecorder()
	auth_controller.Login(w, r)
	if w.Code == http.StatusOK {
		t.Errorf("user was logged in with invalid credentials")
	}
}

func TestLoginWithInValidBody(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\"password\":\"12\"}"))
	w := httptest.NewRecorder()
	auth_controller.Login(w, r)
	if w.Code == http.StatusOK {
		t.Errorf("user was logged in with invalid credentials")
	}
}

func TestLoginWithValidCredentials(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader("{\"username\":\"tom\",\"password\":\"123\"}"))
	w := httptest.NewRecorder()
	auth_controller.Login(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code after successful login. Expected: %d Recieved: %d", http.StatusOK, w.Code)
	}

	token, err := io.ReadAll(w.Result().Body)

	if err != nil {
		t.Errorf("unexpected error parsing data from signup token: %s", err.Error())
	}
	if string(token) == "" {
		t.Errorf("not token returned")
	}
}
