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
	mustKey := []string{"mobile", "code"}
	// 非空判断
	validate := Validate(mustKey, params)
	if len(validate) != 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", validate)))
		return
	}
	mobile := utils.GetParamsStringOfMap(params, mustKey[0])
	if len(mobile) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[0])))
		return
	}
	code := utils.GetParamsStringOfMap(params, mustKey[1])
	if len(code) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[1])))
		return
	}
	// 格式校验
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(mobile, mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	//用户信息
	userinfo, err := selectUserByMobile(option.DB, mobile)
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"mobile": mobile,
		}).Info("用户数据为空")
		return
	}
	// 校验验证码
	verificationKey := fmt.Sprintf("GET_VERIFICATION_CODE_%v", mobile)
	verificationVal := cache.Client.Get(ctx, verificationKey).Val()
	if len(verificationVal) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "验证码失效，请重新获取"))
		return
	}
	if code != verificationVal {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号或验证码错误"))
	}
	// 删除验证码缓存
	cache.Client.Del(ctx, verificationKey)
	// 缓存token
	u := uuid.NewString()
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	cache.Client.Set(ctx, redisKey, u, time.Hour * 12)
	result := make(map[string]string)
	result["token"] = u
	ctx.JSON(http.StatusOK, baseview.GetView(result, ""))
}

// HandlerLoginPassword 手机号 密码登录
func HandlerLoginPassword(ctx *gin.Context) {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	mustKey := []string{"mobile", "password"}
	// 非空校验
	validate := Validate(mustKey, params)
	if len(validate) != 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", validate)))
		return
	}
	mobile := utils.GetParamsStringOfMap(params, mustKey[0])
	if len(mobile) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[0])))
		return
	}
	password := utils.GetParamsStringOfMap(params, mustKey[1])
	if len(password) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[1])))
		return
	}
	// 格式校验
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(mobile, mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	//用户信息
	userinfo, err := selectUserByMobile(option.DB, mobile)
	if err != nil || userinfo == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{
			"mobile": mobile,
		}).Info("用户数据为空")
		return
	}
	// 缓存token
	u := uuid.NewString()
	redisKey := fmt.Sprintf("LOGGING_STATUES_%v", userinfo.Id)
	cache.Client.Set(ctx, redisKey, u, time.Hour * 12)
	// 比较信息是否匹配
	encodePassword := encode.MD5Encode(password, nil)
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
	mustKey := []string{"mobile", "password", "code"}
	validate := Validate(mustKey, params)
	if len(validate) != 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", validate)))
		return
	}
	mobile := utils.GetParamsStringOfMap(params, mustKey[0])
	if len(mobile) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[0])))
		return
	}
	password := utils.GetParamsStringOfMap(params, mustKey[1])
	if len(password) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[1])))
		return
	}
	code := utils.GetParamsStringOfMap(params, mustKey[2])
	if len(code) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[2])))
		return
	}
	// 格式校验
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(mobile, mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "参数错误：手机号格式不正确"))
		return
	}
	// 数据去重
	if isExistMobile(option.DB, mobile) {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号重复"))
		return
	}
	// 校验验证码
	verificationKey := fmt.Sprintf("GET_VERIFICATION_CODE_%v", mobile)
	verificationVal := cache.Client.Get(ctx, verificationKey).Val()
	if len(verificationVal) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "验证码失效，请重新获取"))
		return
	}
	if code != verificationVal {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "验证码不正确"))
		return
	}
	ip := utils.DefaultClientIP(ctx.Request)
	nowTime := time.Now().Unix()
	encodePassword := encode.MD5Encode(password, nil)
	// 准备
	user := &User{
		Mobile: mobile,
		LastLoginIp: ip,
		LastLoginTime: int(nowTime),
		CreateTime: int(nowTime),
		UserPass: encodePassword,
		UserStatus: int8(1),
	}
	// 注册
	userId, err := registryUser(option.DB, user)
	if err != nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "用户注册失败"))
		return
	}
	// 删除验证码缓存
	cache.Client.Del(ctx, verificationKey)
	result := make(map[string]int)
	result["id"] = userId
	ctx.JSON(http.StatusOK, baseview.GetView(result, ""))
}

// HandlerVerification 获取短信验证码
func HandlerVerification(ctx *gin.Context)  {
	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)
	// 非空判断
	mustKey := []string{"mobile"}
	validate := Validate(mustKey, params)
	if len(validate) != 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", validate)))
		return
	}
	mobile := utils.GetParamsStringOfMap(params, mustKey[0])
	if len(mobile) == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, fmt.Sprintf("参数错误：%v为空", mustKey[0])))
		return
	}
	mobileFormat := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	regex := utils.PatternRegex(mobile, mobileFormat)
	if !regex {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "手机号格式不正确"))
		return
	}
	redisKey := fmt.Sprintf("GET_VERIFICATION_CODE_%v", mobile)
	timeoutKey := fmt.Sprintf("GET_VERIFICATION_CODE_TIMEOUT_%v", mobile)
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

func Validate(mustKey []string, params map[string]interface{}) string {
	if len(mustKey) == 0 || len(params) == 0 {
		return ""
	}
	var ok bool
	for _, v := range mustKey{
		if _, ok = params[v]; !ok {  // map 中 key 不存在
			return v
		}
	}
	return ""
}