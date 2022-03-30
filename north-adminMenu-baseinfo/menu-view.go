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

//防止表加s 123
func (AdminMenu) TableName() string {
	return "admin_menu"
}
