package north_role_baseinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	sql_operation "north-project/north-common/sql-operation"
	"strings"
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


//
func RoleDel(id string)  (err error){
	err = sql_operation.GDB.Where("id = ?", id).Delete(&Role{}).Error
	return
}





func InsertSql(tableName string, s1 map[string]interface{}) string {
	columns := ""
	values := ""
	for k, v := range s1 {
		columns = columns + "`" + k + "`,"
		values = values + "'" + v.(string) + "',"
	}
	columns = strings.TrimRight(columns, ",")
	values = strings.TrimRight(values, ",")
	sql := "INSERT INTO `" + tableName + "` " + "(" + columns + ") VALUES (" + values + ")"
	return sql
}
