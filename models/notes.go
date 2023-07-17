package models

type Note struct {
	Id   int32  `json:"id,omitempty"`
	Note string `json:"note,omitempty"`
	SID  string `json:"s_id,omitempty"`
}

type NoteList struct {
	Notes []*Note `json:"notes"`
}
