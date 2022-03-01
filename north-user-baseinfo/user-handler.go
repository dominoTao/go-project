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
// HandlerLoginVerification 手机号 验证码登录
func HandlerLoginVerification(ctx *gin.Context)  {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	// 非空判断
	if !utils.ValidateString(params["mobile"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号为空"))
		return
	}
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(params["mobile"].(string), mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	code := params["code"].(string)
	if !utils.ValidateString(code) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：验证码为空"))
		return
	}
	//用户信息
	userinfo, err := selectUserByMobile(option.DB, params["mobile"].(string))
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"mobile": params["mobile"].(string),
		}).Info("用户数据为空")
		return
	}
	verifKey := fmt.Sprintf("GET_VERIFICATION_CODE_%v", params["mobile"].(string))
	verifVal := cache.Client.Get(ctx, verifKey).Val()
	if len(verifVal) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "验证码失效，请重新获取"))
		return
	}
	// 缓存token
	u := uuid.NewString()
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	cache.Client.Set(ctx, redisKey, u, time.Hour * 12)
	if len(verifVal) != 0 && code == verifVal {
		result := make(map[string]string)
		result["token"] = u
		ctx.JSON(http.StatusOK, baseview.GetView(result, ""))
	}else {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号或验证码错误"))
	}
}

// HandlerLoginPassword 手机号 密码登录
func HandlerLoginPassword(ctx *gin.Context) {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	// 非空判断
	if !utils.ValidateString(params["mobile"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号为空"))
		return
	}
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(params["mobile"].(string), mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	//用户信息
	userinfo, err := selectUserByMobile(option.DB, params["mobile"].(string))
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"mobile": params["mobile"].(string),
		}).Info("用户数据为空")
		return
	}

	// 缓存token
	u := uuid.NewString()
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	cache.Client.Set(ctx, redisKey, u, time.Hour * 12)

	encodePassword := encode.MD5Encode(params["password"].(string), nil)
	// 比较信息是否匹配
	if encodePassword == userinfo.UserPass {
		result := make(map[string]string)
		result["token"] = u
		ctx.JSON(http.StatusOK, baseview.GetView(result, ""))
	}else {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "账号或密码错误"))
	}
}

// HandlerRegistry user registry
func HandlerRegistry(ctx *gin.Context) {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	// 非空判断
	if !utils.ValidateString(params["mobile"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号为空"))
		return
	}
	if !utils.ValidateString(params["password"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：密码为空"))
		return
	}
	// 准备sql参数
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(params["mobile"].(string), mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	if !utils.ValidateString(params["code"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：验证码为空"))
		return
	}

	ip := utils.DefaultClientIP(ctx.Request)
	nowTime := time.Now().Unix()
	encodePassword := encode.MD5Encode(params["password"].(string), nil)
	// 数据去重
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

// HandlerVerification 获取短信验证码
func HandlerVerification(ctx *gin.Context)  {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	if !utils.ValidateString(params["mobile"].(string)) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号为空"))
		return
	}
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(params["mobile"].(string), mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号格式不正确"))
		return
	}
	redisKey := fmt.Sprintf("GET_VERIFICATION_CODE_%v", params["mobile"].(string))
	timeoutKey := fmt.Sprintf("GET_VERIFICATION_CODE_TIMEOUT_%v", params["mobile"].(string))
	val := cache.Client.Get(ctx, redisKey).Val()
	if len(val) != 0 {
		b, _ := cache.Client.Get(ctx, timeoutKey).Bool()
		if b {
			ctx.JSON(http.StatusOK, baseview.GetView(nil, "操作频繁，请稍后再试"))
			return
		}
	}

	sms := utils.GetSMS()
	cache.Client.Set(ctx, redisKey, sms, time.Minute*5)
	cache.Client.Set(ctx, timeoutKey, true, time.Minute)
	ctx.JSON(http.StatusOK, baseview.GetView(sms, ""))
}


