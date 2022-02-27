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
	Id            int64  `json:"id,omitempty"`
	Sex           int8   `json:"sex,omitempty"`
	Birthday      int    `json:"birthday,omitempty"`
	LastLoginTime int    `json:"last_login_time,omitempty"`
	CreateTime    int    `json:"create_time,omitempty"`
	UserStatus    int8   `json:"user_status,omitempty" default:"1"`
	UserLogin     string `json:"user_login,omitempty"`
	UserPass      string `json:"-"`
	UserNickname  string `json:"user_nickname,omitempty"`
	UserEmail     string `json:"user_email,omitempty"`
	UserUrl       string `json:"user_url,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Signature     string `json:"signature,omitempty"`
	LastLoginIp   string `json:"last_login_ip,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
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
	err := DB.QueryRow("SELECT id,user_login,user_pass FROM user WHERE user_login =  ?", username).Scan(&pe.Id, &pe.UserLogin, &pe.UserPass)
	//DB.Query("SELECT id,user_login,user_pass FROM user WHERE user_login =  ?")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	marshal, _ := json.Marshal(pe)
	fmt.Println(string(marshal))
	return &pe, nil
}
