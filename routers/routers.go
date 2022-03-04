package routers

import (
	"github.com/gin-gonic/gin"
	adminMenu "north-project/north-adminMenu-baseinfo"
	role "north-project/north-role-baseinfo"
	user "north-project/north-user-baseinfo"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
}
func SetupRouters() *gin.Engine {

	// user package
	r.Handle("POST", "/login", user.HandlerLoginVerification)
	r.Handle("POST", "/loginP", user.HandlerLoginPassword)
	r.Handle("POST", "/registry", user.HandlerRegistry)
	// 获取验证码
	r.Handle("POST", "/verCode", user.HandlerVerification)

	// role package
	r.Handle("POST", "/role", role.HandlerRoles)
	r.Handle("POST", "/roleAdd", role.HaddlerRoleAdd)
	//r.Handle("POST", "/roleDelete", role.HaddlerRoleDelete) //角色删除
	r.Handle("GET", "/test",role.Test)

	//菜单接口
	r.Handle("GET", "/adminMenuList",adminMenu.HandlerAdminMenuList)
	return r
}



