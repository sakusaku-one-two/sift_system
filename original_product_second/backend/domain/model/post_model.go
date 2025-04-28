package model

import (
	"gorm.io/gorm"
)


type Post struct {
	Id string `gorm:"primarykey" json:"id"`
	
	PostName string `gorm:"type:varchar(100);not null" json:"postName"`
	PostShortName string `gorm:"type:varchar(100);not null" json:"postShortName"`
	Client Client `gorm:"foreginkey:Id" json:"client"`




}


func NewPost(post_name string,client_id string) (Post,error) {

}

