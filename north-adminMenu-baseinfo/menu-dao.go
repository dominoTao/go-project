package north_adminMenu_baseinfo

import sql_operation "north-project/north-common/sql-operation"

//获取菜单列表
func GetAllMenu() (menuList []*AdminMenu, err error) {
	if err := sql_operation.GDB.Where(" status = 1 ").Find(&menuList).Error; err != nil {
		return nil, err
	}
	return
}

//根据id查找菜单
func getById(id int) (menu []*AdminMenu, err error) {
	if err := sql_operation.GDB.Where(" id = ?  AND status = ? ", id, 1).Find(&menu).Error; err != nil {
		return nil, err
	}
	return menu,nil
}


//插入数据
func menuInsert(menuList AdminMenuAdd) (id int,err error) {
	if err := sql_operation.GDB.Create(&menuList).Error; err != nil {
		return 0, err
	}
	return menuList.Id, nil
}

//删除数据