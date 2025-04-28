package util

import (
	"testing"
	"reflect"
)


func Test_int(t *testing.T) {
	Print("数値変換の確認テスト")

	const target = "123"
	result := ToInt(target)
	type_name := reflect.TypeOf(result)
	Print(result,type_name)
}