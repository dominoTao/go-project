package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
	option "north-project/north-common/sql-operation"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
}
func SetupRouters() *gin.Engine {
	r.Handle("POST", "/login", handlerLogin)
	return r
}

func handlerLogin(ctx *gin.Context) {
	db, err := option.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := baseview.BaseResponse{}
	response.Message = "Successful"
	response.Code = 1
	//response.Data = users
	//users := option.SelectUsers(db)
	flag := option.SelectUserByName(db, "张三", "123ewq")
	if flag {
		response.Data = "登录成功"
	} else {
		response.Data = "登录失败"
	}
	ctx.JSON(http.StatusOK, response)
}