package north_user_baseinfo

import (
	"database/sql"
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
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &pe, nil
}

func selectUserByMobile(DB *sql.DB, mobile string) (*User, error) {
	var pe User
	err := DB.QueryRow("SELECT id,mobile,user_pass FROM user WHERE user_login =  ?", mobile).Scan(&pe.Id, &pe.Mobile, &pe.UserPass)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &pe, nil
}

func isExistMobile(DB *sql.DB, mobile string) bool {
	var df int
	err := DB.QueryRow("SELECT id FROM user WHERE mobile=?", mobile).Scan(&df)
	if err != nil || df == 0 {
		return false
	}
	return true
}

func isExistUsername(DB *sql.DB, username string) bool {
	var df int
	err := DB.QueryRow("SELECT id FROM user WHERE user_login=?", username).Scan(&df)
	if err != nil || df == 0 {
		return false
	}
	return true
}

func registryUser(DB *sql.DB, user *User) (int, error) {
	//exec, err := DB.Exec("INSERT INTO user (mobile, last_login_ip, user_login) VALUES (?,?,?)", user.Mobile,user.LastLoginIp,user.UserLogin)
	exec, err := DB.Exec("INSERT INTO user (mobile, last_login_ip, last_login_time, create_time, user_pass, user_status) VALUES (?,?,?,?,?,?)", user.Mobile,user.LastLoginIp, user.LastLoginTime,user.CreateTime, user.UserPass, user.UserStatus)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return int(id), nil
}