package routers

import (
	"github.com/gin-gonic/gin"
	"north-project/north-user-baseinfo"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
}
func SetupRouters() *gin.Engine {
	r.Handle("POST", "/login", north_user_baseinfo.HandlerLogin)
	return r
}

