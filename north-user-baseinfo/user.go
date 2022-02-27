package north_user_baseinfo

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

/**
用户表对应结构体
*/
type User struct {
	Id            int64  `json:"-"`
	Sex           int8   `json:"sex,string"`
	Birthday      int    `json:"birthday"`
	LastLoginTime int    `json:"last_login_time"`
	CreateTime    int    `json:"create_time"`
	UserStatus    int8   `json:"user_status" default:"1"`
	UserLogin     string `json:"user_login"`
	UserPass      string `json:"user_pass"`
	UserNickname  string `json:"user_nickname"`
	UserEmail     string `json:"user_email"`
	UserUrl       string `json:"user_url"`
	Avatar        string `json:"avatar"`
	Signature     string `json:"signature"`
	LastLoginIp   string `json:"last_login_ip"`
	Mobile        string `json:"mobile"`
	More          string `json:"more,omitempty"`
}

func SelectUserById(DB *sql.DB, id int) User {
	var pe User
	err := DB.QueryRow("SELECT id,user_login,user_pass FROM user WHERE id = ?", id).Scan(&pe.Id, &pe.UserLogin, &pe.UserPass)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return pe
}

func SelectUserByUserName(DB *sql.DB, username string) (*User, error) {
	var pe User
	err := DB.QueryRow("SELECT id,user_login,user_pass FROM user WHERE user_login =  ? ", username).Scan(&pe.Id, &pe.UserLogin, &pe.UserPass)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	marshal, _ := json.Marshal(pe)
	fmt.Println(string(marshal))
	return &pe, nil
}
