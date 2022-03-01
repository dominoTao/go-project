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
	Name       string `json:"name,omitempty" binding:"required"`
	Remark     string `json:"remark,omitempty"`
}