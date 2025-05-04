package server

import (
	"backend/util"
	"fmt"
	"reflect"
	"testing"
)

func Test_Is_serverType(t *testing.T) {
	e := CreateServer()
	type_name := reflect.TypeOf(e)

	fmt.Println(type_name)

	fmt.Println(util.TypeOfstring(e))
}
