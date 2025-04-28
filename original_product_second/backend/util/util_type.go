package util

import (
	"reflect"
)


func TypeOfstring(target any) string {
	return 	reflect.TypeOf(target).String()
}

func TypeOf(target any) reflect.Type {
	return reflect.TypeOf(target)
}

func IsStruct(target any) bool {
	type_ := TypeOf(target)
	if  type_.Kind() != reflect.Struct {
		return false
	}
	return true
}

func TypeEq[Model any](target any) bool {
	_, ok := target.(Model)
	return ok
}

