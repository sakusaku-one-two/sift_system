package dto 


import (
	"encoding/json"
)



func ModelToJson[M any](target *M) ([]byte,error) {
	return json.Marshal(target)	
}



func ModleToJsonString[M any](target *M)(string,error) {
	completed,err := ModelToJson[M](target)
	if err != nil {
		return "",err
	}
	string_json := string(completed)
	return string_json,nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~[JSON => MODLE]~~~~~~~~~~~~~~~~~~~~~~~

func JsonToModle[M any](json_byte []byte)(*M,error) {
	Result := new(M) 
	if err := json.Unmarshal(json_byte,Result); err != nil {
		return nil,err
	}

	return Result,nil
}