package north_user_baseinfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
	sql_operation "north-project/north-common/sql-operation"
)

func HandlerLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	//password := ctx.PostForm("passwword")
	fmt.Println("入参 username ：", username)
	//用户信息
	userinfo, err := SelectUserByUserName(sql_operation.DB, username)
	var view *baseview.BaseResponse
	if err != nil {
		view = baseview.GetView(nil, err.Error())
	} else {
		view = baseview.GetView(userinfo, "")
	}
	ctx.JSON(http.StatusOK, view)
}
