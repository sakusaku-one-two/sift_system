package model

import (
	"gorm.io/gorm/"
)



type Employee struct {
	ID uint `gorm:"primaryKey;auto incriment" json:"id"`
	BijiconID uint `gorm:"primarykey" json:"bijiconId"`

	Name string `gorm:"text;not null" json:"name"`
	Kana string `gorm:"text" json:"kana"`
	
	IsManager bool `gorm:"boolian;defalt false" json:"isManager"`
	GrantLevel uint `gorm:"default 0" json:"grantLevel"`

	Email string `gorm:"text" json:"email"`
	

}