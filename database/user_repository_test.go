package database

import (
	"testing"

	"github.com/idugan100/fitness_server_330/models"
)

func TestGetAllUsers(t *testing.T) {
	res, err := user_repo.All()
	if err != nil {
		t.Errorf("unexpected error when getting all users %s", err.Error())
	}

	if len(res) != 3 {
		t.Errorf("unexpected number of users. expected 3 got %d", len(res))
	}
}

func TestUserValidLogin(t *testing.T) {
	user := models.User{UserName: "tom", Password: "123"}
	retrieved_user, err := user_repo.Login(user)
	if err != nil {
		t.Errorf("unexpected error when logging in with valid credentials %s", err.Error())
	}

	if retrieved_user.UserName != user.UserName || retrieved_user.Password != user.Password {
		t.Errorf("retrieved user does not match actual user")
	}
}
