package north_role_baseinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	sql_operation "north-project/north-common/sql-operation"
	"time"
)

func selectRoleById(DB *sql.DB, id int) (*Role, error) {
	var pe Role
	err := DB.QueryRow("SELECT id,status,list_order,name FROM role WHERE id = ?", id).Scan(&pe.Id, &pe.Status, &pe.ListOrder, &pe.Name)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	marshal, _ := json.Marshal(pe)
	fmt.Println(string(marshal))
	return &pe, nil
}

func selectAllRole(DB *sql.DB) (*[]Role, error) {
	result, err := DB.Query("SELECT id,status,list_order,name FROM role WHERE status = 1")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	roles := make([]Role, 0)
	for result.Next() {
		var role Role
		err := result.Scan(&role.Id, &role.Status, &role.ListOrder, &role.Name)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		roles = append(roles, role)
	}
	fmt.Println(roles)
	return &roles, nil
}

//角色添加
func RoleInsert(DB *sql.DB, status int, order int, name string, remark string) (int, error) {
	now := time.Now()
	r, err := DB.Exec("insert into role(status, create_time, update_time,list_order,name,remark) values (?, ?, ?, ?, ?, ?)", status, now.Unix(), now.Unix(), order, name, remark)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, fmt.Errorf(err.Error())
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, fmt.Errorf(err.Error())
	}
	return int(id), nil
}

//根据主键删除
func RoleDel(id int) (err error) {
	err = sql_operation.GDB.Where("id = ?", id).Delete(&Role{}).Error
	if err != nil {
		fmt.Println("exec failed, ", err)
		return fmt.Errorf(err.Error())
	}
	return
}

//根据主键获取角色信息
func getById(id int) (roleInfo []*Role,  err error) {
	if err := sql_operation.GDB.Where("id = ?", id).Find(&Role{}).Error; err != nil {
		return nil, err
	}
	return roleInfo,nil
}


//根据id 更新数据
func updateData(id int,roleEdit RoleEdit) (err error) {
	if err := sql_operation.GDB.Where(" id = ?  AND status = ? ", id, 1).Find(&roleEdit).Error; err != nil {
		return  err
	}
	return
}
