package controllers

import (
	"net/http"
	"net/http/httptest"
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
