package north_user_baseinfo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/session"
	option "north-project/north-common/sql-operation"
)

func HandlerLogin(ctx *gin.Context) {
	var view *baseview.BaseResponse
	m := make(map[string]interface{})
	ctx.BindJSON(&m)

	//用户信息
	userinfo, err := SelectUserByUserName(option.DB, m["username"].(string))
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, err.Error()))
		return
	}
	// 存入session
	//session.SaveSession(ctx.Writer, ctx.Request)

	// 密码加密
	encodePassword := session.MD5Encode(m["password"].(string), nil)
	// 比较信息是否匹配
	m2 := make(map[string]string)
	m2["token"] = session.MD5Encode(m["username"].(string), nil)
	if encodePassword == userinfo.UserPass {
		view = baseview.GetView(m2, "")
	}else {
		view = baseview.GetView(nil, "用户名或密码错误")
	}
	ctx.JSON(http.StatusOK, view)
}

func convObj2String(input interface{}) string {
	if input != nil {
		if s, ok := input.(string); ok{
			return s
		}
	}
	return ""
}