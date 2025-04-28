package util

import (
	"testing"
)



type SampleStructVer2 struct {
	name string
	age int
}

func Test_type(t *testing.T) {
	const val = 123

	type_name := TypeOf(val)

	Print(type_name)

	sample := new(SampleStructVer2)

	type_name = TypeOf(*sample)

	Print(type_name)



}