package north_adminMenu_baseinfo

import sql_operation "north-project/north-common/sql-operation"

//获取菜单列表
func GetAllMenu() (menuList []*AdminMenu, err error) {

	if err := sql_operation.GDB.Where(" status = 1 ").Find(&menuList).Error; err != nil {
		return nil, err
	}
	return
}
