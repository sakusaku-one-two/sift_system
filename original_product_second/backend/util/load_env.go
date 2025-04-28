package util


import (
	"os"
	"fmt"
)


func SetEnv(env_name , value string) {
	os.Setenv(env_name,value)
}

func GetEnv(env_name string) string {
	return os.Getenv(env_name)
}

func GetEnvOrDefault(env_name string,default_value string) string {
	fetched_value := GetEnv(env_name)
	if fetched_value == "" {
		return default_value
	}
	return fetched_value
}

func GetEnvOrDefaultToInt(env_name,defalut_value string) int {
	fetched_value := GetEnvOrDefault(env_name,defalut_value)
	if fetched_value == "-1" {
		return -1
	}

	result := ToInt(fetched_value)
	if  result == -1 {
		panic(fmt.Errorf("GetEnvOrDefaultToInt関数で対象の値が数値に変換できません:%s",fetched_value))
	}

	return result
}