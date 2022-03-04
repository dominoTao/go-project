package sql_operation

import "github.com/jinzhu/gorm"

var (
	GDB *gorm.DB
)

//定义数据库
func InitMySQL() (err error) {
	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":3306)/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	GDB, err = gorm.Open("mysql", dsn) //这个地方不能用 冒号等于  因为 冒号等于就是在方法内申明   而这个地方 是为了申明一个全局的
	if err != nil {
		return
	}
	return GDB.DB().Ping() //返回一个错误
}

//关闭数据库
func Close() (err error) {
	GDB.Close()
	return
}
