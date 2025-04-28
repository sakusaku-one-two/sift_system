package config


import (
	"backend/util"
)

var( 
	DB_NAME = util.GetEnv("DB_NAME")
	DB_USER = util.GetEnv("DB_USER")
	DB_PASSWORD = util.GetEnv("DB_PASSWORD")

)

