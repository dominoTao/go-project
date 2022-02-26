package routers

import (
	"github.com/gin-gonic/gin"
	"north-project/north-user-baseinfo/pkg/view"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
}
func SetupRouters() *gin.Engine {
	r.Handle("POST", "/login", view.HandlerLogin)
	return r
}

