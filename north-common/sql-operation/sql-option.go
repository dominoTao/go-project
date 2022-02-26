package sql_operation

type SqlOption interface {
	Insert(sql []string)
	Delete()
	Update()
	Query()
}
