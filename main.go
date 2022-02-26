package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
	option "north-project/north-common/sql-operation"
)

func main() {
	route := gin.Default()
	route.SetTrustedProxies([]string{"127.0.0.1"})
	route.Handle("POST", "/login", func(ctx *gin.Context) {
		//route.Handle("GET", "/login", func(ctx *gin.Context) {
		//	request := ctx.Request
		//var bb []byte
		//_, _ = request.Body.Read(bb)
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

		//123
		//test
		//test

		//test
		//test
		ctx.JSON(http.StatusOK, response)
	})
	route.Run("localhost:8080")
}
