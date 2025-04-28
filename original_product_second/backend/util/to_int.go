package util

import (
	"strconv"
)



//失敗したら−１を返す
func ToInt(target string) int {
	val,err := strconv.Atoi(target)
	if err != nil {
		return -1
	}
	return val
}