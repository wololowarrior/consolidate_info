package datastore

import (
	"fmt"
	"net/http"

	"accuknox/models"
	"gopkg.in/mgo.v2/bson"
)

type UserDatastore struct {
	users []*models.User
}

func (u *UserDatastore) Insert(user *models.User) ([]byte, *models.DaoError) {
	for _, tempUser := range u.users {
		if tempUser.Email == user.Email {
			return nil, &models.DaoError{
				Message:    fmt.Sprintf("already exists %s", user.Email),
				HttpStatus: http.StatusConflict,
			}
		}
	}
	user.Id = []byte(bson.NewObjectId())
	u.users = append(u.users, user)
	return user.Id, nil
}

func (u *UserDatastore) Login(user *models.UserLoginRequest) (string, *models.DaoError) {
	//TODO implement me
	for _, tempUser := range u.users {
		if tempUser.Email == user.Email {
			if tempUser.Password == user.Password {
				tempUser.SID = "loggedin"
				return tempUser.SID, nil
			} else {
				return "", &models.DaoError{
					Message:    "Wrong Password",
					HttpStatus: http.StatusUnauthorized,
				}
			}
		}
	}
	return "", &models.DaoError{
		Message:    "No such user",
		HttpStatus: http.StatusUnauthorized,
	}
}

func (u *UserDatastore) Logout(user *models.UserLoginRequest) *models.DaoError {
	//TODO implement me
	for _, tempUser := range u.users {
		if tempUser.Email == user.Email {
			if tempUser.Password == user.Password {
				tempUser.SID = ""
				return nil
			}
		}
	}
	return &models.DaoError{
		Message:    "No such user",
		HttpStatus: http.StatusForbidden,
	}
}

func (u *UserDatastore) IsAuthorised(sid string) *models.DaoError {
	for _, tempUser := range u.users {
		if tempUser.SID == sid {
			return nil
		}
	}
	return &models.DaoError{
		Message:    "Session invalid",
		HttpStatus: http.StatusUnauthorized,
	}
}

func NewUsers() *UserDatastore {
	return &UserDatastore{}
}
