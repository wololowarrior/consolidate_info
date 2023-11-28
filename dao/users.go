package dao

import "accuknox/models"

type Users interface {
	Insert(user *models.User) (int64, *models.DaoError)
	GetPrimary(user *models.User) (*models.User, error)
	UpdatePrimary(user *models.User) error
	Get(id int64) (*models.User, *models.DaoError)
}
