package models

type DaoError struct {
	Message    string
	HttpStatus int
}

type HttpErrorResponse struct {
	Message string `json:"message"`
}
