package north_role_baseinfo

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"north-project/north-common/baseview"
	"north-project/north-common/log"
	option "north-project/north-common/sql-operation"
)

//获取角色列表
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

//角色删除
func HaddlerRoleDelete(ctx *gin.Context)  {

	params := make(map[string]interface{})
	_ = ctx.BindJSON(&params)

	_, ok := params["id"]
	if !ok {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "id 不能为空"))
		return
	}

	role,_ := getById(int(int64(params["id"].(float64))))
	if role == nil{
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "未查询到该角色"))
		return
	}

	err := RoleDel(int(int64(params["id"].(float64))))
	if err != nil  {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, "删除失败"))
		// 记录日志
		log.Logger().WithFields(logrus.Fields{}).Info("删除失败")
		return
	}
	ctx.JSON(http.StatusOK, baseview.GetView(1, ""))
}