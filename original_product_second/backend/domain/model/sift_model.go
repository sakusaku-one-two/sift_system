package model


import (
	"gorm.io/gorm"
)


/*
	誰がどこの現場に何時から何時まで勤務するかのレコード
*/

type Sift struct {
	Id string `gorm:"primarykey;not null" json:"id"`
	BijiconId string `gorm:"text"; json:"bijiconId"`

	Emp Employee `grom:"foreignkey:Id" json:"emp"`
	Location Location `gorm:"foreginkey:Id" json:"location"`
	Post Post `gorm:"foreginkey:Id" json:"post"`

}
