package model

import (
	"gorm.io/gorm"
)


type Clinet struct {
	ID uint `gorm:"primarykey;autoincrement" json:"id"`
	BijiconId uint `gorm:"primarykey;autoincrement" json:"bijiconId"`

	Name string `gorm:"text;not null" json:"name"`
	
}


