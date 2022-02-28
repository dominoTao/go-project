package north_user_baseinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func selectUserById(DB *sql.DB, id int) User {
	var pe User
	err := DB.QueryRow("SELECT id,user_login,user_pass FROM user WHERE id = ?", id).Scan(&pe.Id, &pe.UserLogin, &pe.UserPass)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return pe
}

func selectUserByUserName(DB *sql.DB, username string) (*User, error) {
	var pe User
	err := DB.QueryRow("SELECT id,user_login,user_pass FROM user WHERE user_login =  ?", username).Scan(&pe.Id, &pe.UserLogin, &pe.UserPass)
	//DB.Query("SELECT id,user_login,user_pass FROM user WHERE user_login =  ?")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	marshal, _ := json.Marshal(pe)
	fmt.Println(string(marshal))
	return &pe, nil
}
