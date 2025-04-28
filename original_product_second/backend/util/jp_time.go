package util


import (

	"time"
)

var fixedZone_time = 9*60*60
var jstZone = time.FixedZone("JST",fixedZone_time)



func TO_JpTime(t time.Time) (time.Time,error) {
	jst,err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		return t, err
	}

	return t.In(jst),nil
}


func ToJp_by_offset(t time.Time) time.Time {
	return t.In(jstZone)
}

func To_FormatJst(t time.Time,layout string) (string,error) {
	jstTime,err := TO_JpTime(t)
	if err != nil {
		return "",err
	}
	return jstTime.Format(layout),nil
}