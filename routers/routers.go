package routers

import (
	"github.com/gin-gonic/gin"
	role "north-project/north-role-baseinfo"
	user "north-project/north-user-baseinfo"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	_ = r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
}
func SetupRouters() *gin.Engine {
	r.Handle("POST", "/login", user.HandlerLogin)
	r.Handle("POST", "/role", role.HandlerRoles)
	r.Handle("POST", "/registry", user.HandlerRegistry)
	r.Handle("POST", "/roleAdd", role.RoleAdd)
	return r
}



