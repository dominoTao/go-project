package north_role_baseinfo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/log"
	option "north-project/north-common/sql-operation"
	"time"
)

//获取角色
func HandlerRoles(ctx *gin.Context) {
	//所有角色
	roleList, err := selectAllRole(option.DB)
	if err != nil || roleList == nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "角色数据为空"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{}).Info("角色数据为空")
		return
	}
	ctx.JSON(http.StatusOK, baseview.GetView(roleList, ""))
}

//角色添加
func HaddlerRoleAdd(ctx *gin.Context) {

	// 绑定入参数据到map结构体
	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)

	_, ok := params["status"]
	if !ok {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "status 不能为空"))
		return
	}

	_, ok = params["order"]
	if !ok {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "order 不能为空"))
		return
	}
	_, ok = params["name"]
	if !ok {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "name 不能为空"))
		return
	}


	id, err := RoleInsert(option.DB, int(params["status"].(float64)), int(params["order"].(float64)), params["name"].(string), params["remark"].(string))
	if err != nil || id == 0 {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "添加失败"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{}).Info("添加失败")
		return
	}
	ctx.JSON(http.StatusOK, baseview.GetView(id, ""))
}


//git
func Test(ctx *gin.Context)  {
	now := time.Now()
	scoreMap := make(map[string]interface{}, 8)
	scoreMap["status"] = 1
	scoreMap["create_time"] = now.Unix()
	scoreMap["update_time"] = now.Unix()
	scoreMap["list_order"] = 1
	scoreMap["name"] = "管理员"
	scoreMap["remark"] = "这是一个角色"
	sql  := InsertSql("role",scoreMap)
	ctx.JSON(http.StatusOK, baseview.GetView(sql, ""))
}