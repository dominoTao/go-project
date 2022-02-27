package main

import (
	"fmt"
	"north-project/north-common/log"
	"north-project/routers"
)

func main() {
	router := routers.SetupRouters()
	// 日志写入gin
	router.Use(log.LoggerToFile())
	if err := router.Run(":8070"); err != nil {
		fmt.Errorf("failed start routers")
		return
	}
}
