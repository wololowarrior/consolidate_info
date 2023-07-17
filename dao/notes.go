package dao

import "accuknox/models"

type Notes interface {
	Insert(note *models.Note) (int32, *models.DaoError)
	GetAll(sid string) ([]*models.Note, *models.DaoError)
	Delete(note *models.Note) *models.DaoError
}
