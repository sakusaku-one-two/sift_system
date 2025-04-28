package model

import (
	"testing"
	"backend/dto"
	"backend/util"
)


func emp_Test(t *testing.T) {
	util.Print("start emp test !")

	emp := Employee{

	}

	result := dto.ToJson(emp)

	util.Print(result)

	defer util.Print("emp test done!")
}