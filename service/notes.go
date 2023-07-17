package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"accuknox/dao"
	"accuknox/dao/datastore"
	"accuknox/models"
)

type Notes struct {
	notes dao.Notes
	users dao.Users
}

func NewNotesService(notes *datastore.NotesDatastore, user *datastore.UserDatastore) *Notes {
	return &Notes{notes: notes, users: user}
}

func (n *Notes) Create(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var noteCreate *models.Note
	err = json.Unmarshal(rbody, &noteCreate)
	authorised := n.users.IsAuthorised(noteCreate.SID)
	if authorised != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: "unauthorised"})
		return
	}
	nodeID, daoError := n.notes.Insert(noteCreate)
	if daoError != nil {
		w.WriteHeader(daoError.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoError.Message})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.Note{Id: nodeID})
}
func (n *Notes) List(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var noteCreate *models.Note
	err = json.Unmarshal(rbody, &noteCreate)
	authorised := n.users.IsAuthorised(noteCreate.SID)
	if authorised != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: "unauthorised"})
		return
	}
	notes, daoError := n.notes.GetAll(noteCreate.SID)
	if daoError != nil {
		w.WriteHeader(daoError.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoError.Message})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
func (n *Notes) Delete(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var noteCreate *models.Note
	err = json.Unmarshal(rbody, &noteCreate)
	authorised := n.users.IsAuthorised(noteCreate.SID)
	if authorised != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: "unauthorised"})
		return
	}
	daoError := n.notes.Delete(noteCreate)
	if daoError != nil {
		w.WriteHeader(daoError.HttpStatus)
	}
	w.WriteHeader(http.StatusOK)
}
