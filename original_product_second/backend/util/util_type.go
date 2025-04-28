package util

import (
	"reflect"
)


func TypeOf(target any) string {
	return 	reflect.TypeOf(target).String()
}

