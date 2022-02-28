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
	"north-project/north-common/utils"
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
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	// 非空判断
	if !utils.ValidateString(params["username"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户名为空"))
		return
	}
	if !utils.ValidateString(params["mobile"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号为空"))
		return
	}
	if !utils.ValidateString(params["password"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "密码为空"))
		return
	}
	// 准备sql参数
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(params["mobile"].(string), mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号格式不正确"))
		return
	}
	ip := utils.DefaultClientIP(ctx.Request)
	nowTime := time.Now().Unix()
	encodePassword := encode.MD5Encode(params["password"].(string), nil)
	// 数据去重
	username := isExistUsername(option.DB, params["username"].(string))
	if username {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户名重复"))
		return
	}
	mobile := isExistMobile(option.DB, params["mobile"].(string))
	if mobile {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号重复"))
		return
	}
	// 准备
	user := &User{
		Mobile: params["mobile"].(string),
		LastLoginIp: ip,
		LastLoginTime: int(nowTime),
		CreateTime: int(nowTime),
		UserLogin: params["username"].(string),
		UserPass: encodePassword,
		UserStatus: int8(1),
	}
	userId, err := registryUser(option.DB, user)
	if err != nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户注册失败"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"name": params["username"].(string),
		}).Info("用户注册失败")
		return
	}
	result := make(map[string]int)
	result["id"] = userId
	ctx.JSON(http.StatusOK, baseview.GetView(result, ""))
}

