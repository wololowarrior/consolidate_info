package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"accuknox/dao"
	"accuknox/dao/datastore"
	"accuknox/models"
)

type Users struct {
	user dao.Users
}

func NewUserService(user *datastore.UserDatastore) *Users {
	return &Users{user: user}
}

func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user *models.User
	err = json.Unmarshal(rbody, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(user.Name) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: "invalid name"})
		return
	}
	id, daoErr := u.user.Insert(user)
	if daoErr != nil {
		w.WriteHeader(daoErr.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoErr.Message})
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&models.UserCreatedResponse{Id: id})
}
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var loginRequest *models.UserLoginRequest
	err = json.Unmarshal(rbody, &loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sid, daoError := u.user.Login(loginRequest)
	if daoError != nil {
		w.WriteHeader(daoError.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoError.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&models.UserLoginResponse{SID: sid})
}
func (u *Users) Logout(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var loginRequest *models.UserLoginRequest
	err = json.Unmarshal(rbody, &loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	daoError := u.user.Logout(loginRequest)
	if daoError != nil {
		w.WriteHeader(daoError.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoError.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
}
