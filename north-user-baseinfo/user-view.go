package north_user_baseinfo

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

