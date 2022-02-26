package baseview

type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
	Code int `json:"code"`
	Message string `json:"message"`
}

func getView(s interface{}, msg string) (code, message string, sv interface{}) {
	if s != nil {
		sv = s
		code = "1"
		message = "Successful"
	}else {
		sv = nil
		code = "0"
		message = msg
	}
	return
}

