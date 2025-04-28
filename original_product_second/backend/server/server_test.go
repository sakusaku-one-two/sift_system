package server

import (
	"testing"
	"backend/util"
	"reflect"
	"fmt"
)



func Test_Is_serverType(t *testing.T) {
	e := CreateServer()
	type_name := reflect.TypeOf(e)

	fmt.Println(type_name)

	fmt.Println(util.TypeOfstring(e))
}