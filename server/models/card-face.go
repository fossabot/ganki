package models

import "github.com/jinzhu/gorm"

type CardFace struct {
	gorm.Model
	PrimaryInfo   string
	SecondaryInfo string
}
