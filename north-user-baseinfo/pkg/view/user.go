package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
	option "north-project/north-common/sql-operation"
)

/**
用户表对应结构体
 */
type User struct {
	Id int64 `json:"-"`
	Sex int8 `json:"sex,string"`
	Birthday int `json:"birthday"`
	LastLoginTime int `json:"last_login_time"`
	CreateTime int `json:"create_time"`
	UserStatus int8 `json:"user_status" default:"1"`
	UserLogin string `json:"user_login"`
	UserPass string `json:"user_pass"`
	UserNickname string `json:"user_nickname"`
	UserEmail string `json:"user_email"`
	UserUrl string `json:"user_url"`
	Avatar string `json:"avatar"`
	Signature string `json:"signature"`
	LastLoginIp string `json:"last_login_ip"`
	Mobile string `json:"mobile"`
	More string `json:"more,omitempty"`
}

type UserLogin struct {
	Id int
	Name string
	Pass string
}


func HandlerLogin(ctx *gin.Context) {
	db, err := option.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	flag := option.SelectUserByName(db, "张三", "123ewq")

	response := baseview.BaseResponse{}
	response.Message = "Successful"
	response.Code = 1
	//response.Data = users
	//users := option.SelectUsers(db)
	if flag {
		response.Data = "登录成功"
	} else {
		response.Data = "登录失败"
	}
	ctx.JSON(http.StatusOK, response)
}