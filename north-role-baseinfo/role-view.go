package north_role_baseinfo

/**
角色表映射结构体
*/
type Role struct {
	Id         int64  `json:"id,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime int    `json:"create_time,omitempty"`
	UpdateTime int    `json:"update_time,omitempty"`
	ListOrder  int8   `json:"list_order,omitempty"`
	Name       string `json:"name,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

//防止表加s
func (Role) TableName() string {
	return "role"
}

type RoleEdit struct {
	Id         int64  `json:"id,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime int    `json:"create_time,omitempty"`
	UpdateTime int    `json:"update_time,omitempty"`
	ListOrder  int8   `json:"list_order,omitempty"`
	Name       string `json:"name,omitempty" binding:"required"`
	Remark     string `json:"remark,omitempty"`
}

func (RoleEdit) TableName() string {
	return "role"
}

type RoleEdit struct {
	Id         int64  `json:"id" binding:"required"`
	Status     int    `json:"status" binding:"required"`
	UpdateTime int    `json:"update_time" `
	ListOrder  int8   `json:"list_order"`
	Name       string `json:"name,omitempty" binding:"required"`
	Remark     string `json:"remark,omitempty"`
}

func (RoleEdit) TableName() string {
	return "role"
}
