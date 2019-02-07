package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Picture  string
	Password string
}
