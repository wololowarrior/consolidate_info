package models

type User struct {
	Id             int64   `json:"id"`
	Email          string  `json:"email"`
	Name           string  `json:"name"`
	Phonenumber    int     `json:"phonenumber"`
	Ipv4           string  `json:"ipv4,omitempty"`
	LinkedIds      []int64 `json:"linkedId,omitempty"`
	LinkPrecedence string  `json:"linkPrecedence,omitempty"`
}

type UserCreatedResponse struct {
	Id        int64   `json:"id"`
	LinkedIds []int64 `json:"linkedId,omitempty"`
}

type IdentifyMatchResponse struct {
	PrimaryContactID    int64    `json:"primaryContactID"`
	SecondaryContactIDs []int64  `json:"secondaryContactIDs"`
	Emails              []string `json:"emails"`
	PhoneNumbers        []int    `json:"phoneNumbers"`
}
