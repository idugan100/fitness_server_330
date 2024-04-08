package database

import (
	"database/sql"
	"errors"

	"github.com/idugan100/fitness_server_330/models"
)

type UserRepository struct {
	Connection *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return UserRepository{Connection: conn}
}

func (u UserRepository) Login(user models.User) (models.User, error) {
	rows, err := u.Connection.Query("SELECT * FROM Users WHERE username=? AND password=?", user.UserName, user.Password)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if !rows.Next() {
		return user, errors.New("user Not Found")
	}
	err = rows.Scan(&user.Id, &user.UserName, &user.Password, &user.IsAdmin)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepository) Signup(user models.User) (models.User, error) {
	_, err := u.Connection.Exec("INSERT INTO Users(userName,password) VALUES (?,?)", user.UserName, user.Password)
	if err != nil {
		return user, err
	}
	inserted_user, err := u.Login(user)
	if err != nil {
		return user, err
	}
	return inserted_user, nil
}

func (u UserRepository) All() ([]models.User, error) {
	var userList []models.User
	var user models.User
	rows, err := u.Connection.Query("SELECT * FROM Users")
	if err != nil {
		return userList, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Password, &user.IsAdmin)
		if err != nil {
			return userList, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
