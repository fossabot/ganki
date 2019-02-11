package controllers

import (
	"encoding/json"
	"github.com/dulev/ganki/server/controllers/middleware"
	"github.com/dulev/ganki/server/user"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/dulev/ganki/server/models"
)

// TODO: Implement route handlers

type UserController struct {
	database       *gorm.DB
	sessionManager *middleware.SessionManager
	userService    user.UserService
}

func NewUserController(
	database *gorm.DB,
	sessionManager *middleware.SessionManager,
	userService user.UserService) UserController {
	return UserController{
		database:       database,
		sessionManager: sessionManager,
		userService:    userService,
	}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	uc.sessionManager.ShouldBeLoggedOut(w, r)

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
	uc.sessionManager.ShouldBeLoggedOut(w, r)

	decoder := json.NewDecoder(r.Body)

	type LoginDTO struct {
		Username string
		Password string
	}

	var loginDetails LoginDTO
	err := decoder.Decode(&loginDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = uc.userService.Authenticate(loginDetails.Username, loginDetails.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	uc.sessionManager.Login(loginDetails.Username, w, r)
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	_, err := uc.sessionManager.ShouldBeLoggedIn(w, r)
	if err != nil {
		panic("TODO alabala")
	}

	// Session (Extract)
	uc.sessionManager.Logout(w, r)
}
