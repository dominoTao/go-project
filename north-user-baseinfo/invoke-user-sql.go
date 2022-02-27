package north_user_baseinfo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/log"
	"north-project/north-common/session"
	option "north-project/north-common/sql-operation"
)

func HandlerLogin(ctx *gin.Context) {
	// 绑定入参数据到map结构体
	m := make(map[string]interface{})
	ctx.BindJSON(&m)
	//用户信息
	userinfo, err := SelectUserByUserName(option.DB, m["username"].(string))
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"name": m["username"].(string),
		}).Info("用户数据为空")
		return
	}
	// TODO 存入session
	//session.SaveSession(ctx.Writer, ctx.Request)
	// 密码加密
	//TODO 抽象出断言方法
	encodePassword := session.MD5Encode(m["password"].(string), nil)
	// 比较信息是否匹配
	m2 := make(map[string]string)
	m2["token"] = session.MD5Encode(m["username"].(string), nil)

	var view *baseview.BaseResponse
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