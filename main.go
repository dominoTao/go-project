package main

import (
	"fmt"
	"north-project/north-common/cache"
	"north-project/north-common/log"
	sql_operation "north-project/north-common/sql-operation"
	"north-project/routers"
)

func main() {
	router := routers.SetupRouters()
	// 日志写入gin
	router.Use(log.LoggerToFile())
	// 连接redis
	cache.InitRedis()

	//连接数据库 GROM方式链接
	err := sql_operation.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer sql_operation.Close() //程序退出关闭数据库连接 defer 关键字 表示程序运行结束之后调用的方法

	if err := router.Run(":8080"); err != nil {
		fmt.Errorf("failed start routers")
		return
	}


}
