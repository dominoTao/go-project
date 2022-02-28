package utils

import (
	"github.com/modern-go/reflect2"
)

func ValidateString(i string) bool {
	return len(i) > 0
}
//TODO 待处理 不能用
func ValidateInterface(i interface{}) bool {
	if i == nil {
		return false
	}
	//reflect.DeepEqual()
	return !reflect2.IsNil(i)
}