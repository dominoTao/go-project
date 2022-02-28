package north_user_baseinfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/cache"
	"north-project/north-common/encode"
	"north-project/north-common/log"
	option "north-project/north-common/sql-operation"
	"time"
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

	// 判断token是否失效
	m2 := make(map[string]string)
	m2["token"] = encode.MD5Encode(m["username"].(string), nil)

	// 缓存redis
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	if cache.Client.Exists(ctx, redisKey).Val() != 1 {
		cache.Client.Set(ctx, redisKey, m2["token"], time.Minute * 30)
	}

	//TODO 抽象出断言方法
	// 密码加密
	encodePassword := encode.MD5Encode(m["password"].(string), nil)
	var view *baseview.BaseResponse
	// 比较信息是否匹配
	if encodePassword == userinfo.UserPass {
		view = baseview.GetView(m2, "")
	}else {
		view = baseview.GetView(nil, "用户名或密码错误")
	}
	ctx.JSON(http.StatusOK, view)
}

