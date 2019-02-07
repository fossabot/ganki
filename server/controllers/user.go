package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/dulev/ganki/server/models"
)

// TODO: Implement route handlers

type UserController struct {
	database    *gorm.DB
	userService UserService
}

func NewUserController(database *gorm.DB, userService UserService) UserController {
	return UserController{
		database:    database,
		userService: userService,
	}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var userReg models.User
	err := decoder.Decode(&userReg)
	if err != nil {
		log.Printf("error decoding user registration json: %#v\n", err)
		// TODO: Validate data using govalidator
		w.WriteHeader(http.StatusBadRequest)
	}

	err = uc.userService.RegisterUser(userReg)
	if err != nil {
		// TODO: Do smth
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: NotLoggedInMiddleware

	decoder := json.NewDecoder(r.Body)

	type LoginDTO struct {
		Username string
		Password string
	}

	loginDetails := LoginDTO{}
	err := decoder.Decode(&loginDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = uc.userService.Authenticate(loginDetails.Username, loginDetails.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
