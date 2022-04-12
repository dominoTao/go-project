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
	r.Handle("GET", "/role", role.HandlerRoles)             //角色列表
	r.Handle("POST", "/roleAdd", role.HaddlerRoleAdd)       //角色添加
	r.Handle("POST", "/roleDelete", role.HaddlerRoleDelete) //角色删除
	r.Handle("POST", "/roleEdit", role.HandleRoleEdit)      //角色编辑

	//菜单接口
	r.Handle("GET", "/adminMenuList", adminMenu.HandlerAdminMenuList)
	return r
}
