package north_user_baseinfo

import (
	"fmt"
	"north-project/north-common/utils"
	"testing"
)

func TestGetSMS(t *testing.T) {
	sms := utils.GetSMS()
	fmt.Println(sms)
}

func BenchmarkGetSMS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(utils.GetSMS())
	}
}
