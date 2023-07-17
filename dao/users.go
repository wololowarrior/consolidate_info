package dao

import "accuknox/models"

type Users interface {
	Insert(user *models.User) ([]byte, *models.DaoError)
	Login(user *models.UserLoginRequest) (string, *models.DaoError)
	Logout(user *models.UserLoginRequest) *models.DaoError
	IsAuthorised(sid string) *models.DaoError
}
