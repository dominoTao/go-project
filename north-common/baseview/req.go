package baseview

import "fmt"

type BaseResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func GetView(s interface{}, msg string) *BaseResponse {
	resp := &BaseResponse{}
	if s != nil {
		resp.Code = 1
		resp.Message = "Successful"
		fmt.Println(s)
		resp.Data = s
	} else {
		resp.Code = 0
		resp.Message = msg
		resp.Data = nil
	}
	return resp
}
