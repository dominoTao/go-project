package north_user_baseinfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/cache"
	"north-project/north-common/encode"
	"north-project/north-common/log"
	option "north-project/north-common/sql-operation"
	"time"
)

// HandlerLogin user login
func HandlerLogin(ctx *gin.Context) {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	//用户信息
	userinfo, err := selectUserByUserName(option.DB, params["username"].(string))
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"name": params["username"].(string),
		}).Info("用户数据为空")
		return
	}

	u := uuid.NewString()
	// 缓存redis
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	cache.Client.Set(ctx, redisKey, u, time.Hour * 12)

	encodePassword := encode.MD5Encode(params["password"].(string), nil)
	var view *baseview.BaseResponse
	// 比较信息是否匹配
	if encodePassword == userinfo.UserPass {
		result := make(map[string]string)
		result["token"] = u
		view = baseview.GetView(result, "")
	}else {
		view = baseview.GetView(nil, "用户名或密码错误")
	}
	ctx.JSON(http.StatusOK, view)
}

// HandlerRegistry user registry
func HandlerRegistry(ctx *gin.Context) {

}