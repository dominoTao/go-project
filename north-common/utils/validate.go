package utils

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func ValidateString(i string) bool {
	return len(i) > 0
}
//TODO 待处理 不能用  nil->false;not nil->true
func ValidateInterface(i interface{}) bool {
	if i == nil {
		return false
	}
	//reflect.DeepEqual()
	return !reflect2.IsNil(i)
}

// GetSMS 获取短信验证码
func GetSMS() string {
	var mu sync.Mutex
	mu.Lock()
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000000)
	msg := strconv.Itoa(i)
	if len(msg) < 6 {
		msg = fmt.Sprintf("%s%s", "0", msg)
	}
	mu.Unlock()
	return msg
}
