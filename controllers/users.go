package controllers

import (
	"fmt"
	"net/http"

	"github.com/leondore/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUser{}
	newUser.Email = r.FormValue("email")
	newUser.Password = r.FormValue("password")

	savedUser, err := u.UserService.Create(newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User successfully created: %s", savedUser.Email)
}
