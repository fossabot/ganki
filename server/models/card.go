package models

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Front   CardFace `gorm:"foreignkey:FrontID"`
	FrontID uint
	Back    CardFace `gorm:"foreignkey:BackID"`
	BackID  uint
	DeckID  uint `gorm:"foreign"`
}
