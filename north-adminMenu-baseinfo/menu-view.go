package north_adminMenu_baseinfo

type AdminMenu struct {
	Id        int          `json:"id,omitempty"`
	ParentId  int          `json:"parent_id"`
	Type      int          `json:"type,omitempty"`
	Status    int          `json:"status,omitempty"`
	Action    string       `json:"action,omitempty"`
	ListOrder string       `json:"list_order,omitempty"`
	Name      string       `json:"name,omitempty"`
	Icon      string       `json:"icon,omitempty"`
	Remark    string       `json:"remark,omitempty"`
	Child     []*AdminMenu `json:"child,omitempty"`
}

//防止表加s
func (AdminMenu) TableName() string {
	return "admin_menu"
}

//添加菜单
type AdminMenuAdd struct {
	Id        int    `json:"id" `
	ParentId  int    `json:"parent_id" binding:"required"`
	Action    string `json:"action" binding:"required"`
	ListOrder string `json:"list_order" binding:"required" default:"1"`
	Name      string `json:"name" binding:"required"`
	Icon      string `json:"icon"`
	Remark    string `json:"remark"`
}

func (AdminMenuAdd) TableName() string {
	return "admin_menu"
}

//update
type AdminMenuUpdate struct {
	Id        int    `json:"id,omitempty" binding:"required"`
	ParentId  int    `json:"parent_id" binding:"required"`
	Action    string `json:"action,omitempty" binding:"required"`
	ListOrder string `json:"list_order,omitempty" binding:"required"`
	Name      string `json:"name,omitempty" binding:"required"`
	Icon      string `json:"icon,omitempty"`
	Remark    string `json:"remark,omitempty"`
}
