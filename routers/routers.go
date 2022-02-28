package routers

import (
	"github.com/gin-gonic/gin"
	role "north-project/north-role-baseinfo"
	userinfo "north-project/north-user-baseinfo"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
}
func SetupRouters() *gin.Engine {
	r.Handle("POST", "/login", userinfo.HandlerLogin)
	r.Handle("GET", "/roleList", role.RoleList)
	//r.Handle("POST", "/roleAdd", role.RoleAdd)
	return r
}



