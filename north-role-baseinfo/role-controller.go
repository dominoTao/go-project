package north_role_baseinfo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/log"
	option "north-project/north-common/sql-operation"
)

//获取绝色
func RoleList(ctx *gin.Context) {
	//所有角色
	roleList, err := SelectAllRole(option.DB)
	if err != nil || roleList == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "角色数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{}).Info("角色数据为空")
		return
	}
	ctx.JSON(http.StatusOK, baseview.GetView(roleList, ""))
}

//func RoleAdd(ctx *gin.Context)  {
//	err := RoleAdd(option.DB)
//
//}
