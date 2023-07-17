package models

type User struct {
	Id       []byte
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	SID      string
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	SID string `json:"s_id"`
}

type UserCreatedResponse struct {
	Id []byte `json:"id"`
}
