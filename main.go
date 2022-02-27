package main

import (
	"fmt"
	"north-project/routers"
)

func main() {
	r := routers.SetupRouters()
	if err := r.Run(":8070"); err != nil {
		fmt.Errorf("failed start routers")
		return
	}
}
