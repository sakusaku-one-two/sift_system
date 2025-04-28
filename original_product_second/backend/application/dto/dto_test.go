package dto 

import (
	"backend/util"
	"testing"
)


type SampleModle struct {
	Id  uint `json:"id"`
	Name string `json:"name"`
}


func Test_json(t *testing.T) {
	temp := &SampleModle{
		Id:1,
		Name:"user name",
	}
	util.Print("test json")

	json_encoded,err := ModelToJson[SampleModle](temp)

	util.Print(err,json_encoded)
	util.Print(ModleToJsonString[SampleModle](temp))
	modle_data,err := JsonToModle[SampleModle](json_encoded)
	util.Print("JSON BYTE => MODLE",modle_data.Id)
}