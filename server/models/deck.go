package models

import "github.com/jinzhu/gorm"

type Deck struct {
	gorm.Model
	UserID uint
	Name   string
	Cards  []Card
}
