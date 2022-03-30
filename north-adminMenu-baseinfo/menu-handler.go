package north_adminMenu_baseinfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"north-project/north-common/baseview"
)

var Jsondata []byte

//菜单列表
func HandlerAdminMenuList(ctx *gin.Context) {

	list := make([]*AdminMenu, 0)
	menuList, err := GetAllMenu()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	//拼一个三级的树
	for _, v := range menuList {
		if v.ParentId == 0 {
			//一级
			list = append(list, v)
		}
	}

	for _, v := range list {
		for _, vv := range menuList {
			if v.Id == vv.ParentId {
				v.Child = append(v.Child, vv)
			}
		}
	}

	if err != nil {
		ctx.JSON(http.StatusOK, baseview.GetView(nil, err.Error()))
	} else {
		ctx.JSON(http.StatusOK, baseview.GetView(list, ""))
	}
}

func makeTree(Allnode []*AdminMenu, node *AdminMenu) { //参数为父节点，添加父节点的子节点指针切片
	childs, _ := haveChild(Allnode, node) //判断节点是否有子节点并返回

	if childs != nil {
		fmt.Printf("\n")
		fmt.Println(*node)
		fmt.Println("子节点：")

		for _, v := range childs {
			fmt.Println(*v)
		}

		node.Child = append(node.Child, childs[0:]...)
		for _, v := range childs {
			_, has := haveChild(Allnode, v)
			if has {
				makeTree(Allnode, v)
			}
		}
	}
}

//判断当前节点是否含有子节点
func haveChild(Allnode []*AdminMenu, node *AdminMenu) (childs []*AdminMenu, yes bool) {
	for _, v := range Allnode {
		if v.ParentId == node.Id {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}

func printTree(tree_node *AdminMenu) {
	for _, v := range tree_node.Child {
		fmt.Printf("%d,%d,%s", v.Id, v.ParentId, v.Name)
		fmt.Println("##########")
		if len(v.Child) != 0 {
			printTree(v)
		}
	}
}
