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

	username := ctx.PostForm("username")
	password := ctx.PostForm("passwword")
	//用户信息
	userinfo, err := SelectUserByUserName(option.DB, username)
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, err.Error()))
		return
	}
	// 存入session
	getSess := session.GetSession(ctx.Writer, ctx.Request)
	// 密码加密
	encodePassword := session.MD5Encode(password+convObj2String(getSess), nil)
	// 比较信息是否匹配
	encodeUserPass := session.MD5Encode(userinfo.UserPass, nil)
	if encodePassword == encodeUserPass {
		view = baseview.GetView(userinfo, "")
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