package datastore

import (
	"net/http"

	"accuknox/models"
)

type NotesDatastore struct {
	notes  []*models.Note
	noteID int32
}

func NewNotes() *NotesDatastore {
	return &NotesDatastore{noteID: 1}
}

func (n *NotesDatastore) Insert(note *models.Note) (int32, *models.DaoError) {
	note.Id = n.noteID
	n.noteID++
	n.notes = append(n.notes, note)
	return note.Id, nil
}

func (n *NotesDatastore) GetAll(sid string) ([]*models.Note, *models.DaoError) {
	//TODO implement me
	var noteList []*models.Note
	for _, note := range n.notes {
		if note.SID == sid {
			noteList = append(noteList, &models.Note{
				Note: note.Note,
				Id:   note.Id,
			})
		}
	}
	return noteList, nil
}

func (n *NotesDatastore) Delete(note *models.Note) *models.DaoError {
	//TODO implement me
	for idx, note_1 := range n.notes {
		if note_1.Id == note.Id {
			n.notes = append(n.notes[:idx], n.notes[idx+1:]...)
			return nil
		}
	}
	return &models.DaoError{HttpStatus: http.StatusNotFound}
}
