package datastore

import (
	"database/sql"
	"net/http"

	"accuknox/dao/datastore/queries"
	"accuknox/models"
	"github.com/lib/pq"
)

type UserDatastore struct {
	db *sql.DB
}

const pqErrUniqueViolationName = "unique_violation"
const iDNameUniqueIndex = "email_phone_number_idx"

func NewUsers(db *sql.DB) *UserDatastore {
	return &UserDatastore{db: db}
}

func (u *UserDatastore) Insert(user *models.User) (int64, *models.DaoError) {
	var id int64
	var state = "primary"
	primary, err1 := u.GetPrimary(user)
	if err1 != sql.ErrNoRows {
		state = "secondary"
	}
	err := u.db.QueryRow(
		queries.Insert,
		user.Phonenumber,
		user.Email,
		state,
		user.Ipv4,
	).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == pqErrUniqueViolationName {
				return 0, &models.DaoError{Message: "Already Exists", HttpStatus: http.StatusConflict}
			}
		}
		return 0, &models.DaoError{Message: err.Error(), HttpStatus: http.StatusInternalServerError}
	}

	if err1 != sql.ErrNoRows {
		primary.LinkedIds = append(primary.LinkedIds, id)
		err := u.UpdatePrimary(primary)
		if err != nil {
			if err != sql.ErrNoRows {
				return 0, &models.DaoError{Message: err.Error(), HttpStatus: http.StatusInternalServerError}
			}
		}
	}

	return id, nil
}

func (u *UserDatastore) GetPrimary(user *models.User) (*models.User, error) {
	primaryUser := &models.User{}
	var numbers pq.Int64Array
	err := u.db.QueryRow(
		queries.GetPrimary,
		user.Ipv4,
	).Scan(&primaryUser.Id,
		&primaryUser.Email,
		&primaryUser.Phonenumber,
		&numbers,
		&primaryUser.LinkPrecedence)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	primaryUser.LinkedIds = numbers
	return primaryUser, nil
}

func (u *UserDatastore) UpdatePrimary(user *models.User) error {
	var id int64
	err := u.db.QueryRow(
		queries.UpdateLinkedID,
		pq.Int64Array(user.LinkedIds),
		user.Id,
	).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
	}
	return nil
}

func (u *UserDatastore) Get(id int64) (*models.User, *models.DaoError) {
	user := &models.User{}
	var numbers pq.Int64Array
	err := u.db.QueryRow(
		queries.Get,
		id,
	).Scan(&user.Id,
		&user.Email,
		&user.Phonenumber,
		&numbers,
		&user.Ipv4,
		&user.LinkPrecedence,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.DaoError{Message: "Not found", HttpStatus: http.StatusNotFound}
		}
	}
	user.LinkedIds = numbers
	return user, nil
}
