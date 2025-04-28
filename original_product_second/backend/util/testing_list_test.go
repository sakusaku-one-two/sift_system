
package util

import (
	"testing"
	"fmt"
	"time"
)


type sample struct {
	id int
	name string
}


func TestDelete(t *testing.T) {
	target := &sample{
		id:2,
		name:"name 2",
	}

	l := NewList[sample,int]()
	l.Append(&sample{
		id:1,
		name:"name1",
	})
	l.Append(target)

	

	l.ForEach(func(target *sample){
		Print(target)
	})

	l.Delete(target)
	fmt.Println("deleted to",l.Len())
}


func Test_jp_TIME(t *testing.T) {
	Print("test jp time start!!")
	current_time := time.Now()
	Print("現在の時刻",current_time)

	jp_time,err := TO_JpTime(current_time)
	if err != nil {
		return
	}
	Print("日本時間",jp_time)

}