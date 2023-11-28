package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"accuknox/dao"
	"accuknox/dao/datastore"
	"accuknox/models"
)

type Users struct {
	user dao.Users
}

var covertResponses = []string{
	"Password not upto the strength, make it stronger",
	"Check your emailID",
	"Clear Browser Cache",
	"Browser not supported",
}

func NewUserService(user *datastore.UserDatastore) *Users {
	return &Users{user: user}
}

func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user *models.User
	err = json.Unmarshal(rbody, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(strings.Trim(user.Name, " ")) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: "invalid name"})
		return
	}
	if len(strings.Trim(user.Email, " ")) < 3 {
		http.Error(w, "invalid Email", http.StatusBadRequest)
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// Handle the error, if any
		fmt.Println("Error extracting IP:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Ipv4 = ip

	id, daoErr := u.user.Insert(user)
	if daoErr != nil && daoErr.HttpStatus == http.StatusConflict {
		w.WriteHeader(daoErr.HttpStatus)
		json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: daoErr.Message})
		return
	}

	primary, _ := u.user.GetPrimary(user)
	fmt.Printf("%s\n", primary)

	if primary != nil {
		resp1 := models.IdentifyMatchResponse{
			PrimaryContactID: primary.Id,
		}
		for _, secId := range primary.LinkedIds {
			secUser, err := u.user.Get(int64(secId))
			if err != nil {
				fmt.Println(secId, err)
				continue
			}
			fmt.Printf("sec user %s\n", secUser)
			resp1.Emails = append(resp1.Emails, secUser.Email)
			resp1.SecondaryContactIDs = append(resp1.SecondaryContactIDs, secUser.Id)
			resp1.PhoneNumbers = append(resp1.PhoneNumbers, secUser.Phonenumber)
		}

		//randomIndex := rand.Intn(len(covertResponses))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//err = json.NewEncoder(w).Encode(&models.HttpErrorResponse{Message: covertResponses[randomIndex]})
		err = json.NewEncoder(w).Encode(resp1)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&models.UserCreatedResponse{Id: id})
}
