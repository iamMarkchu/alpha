package main

import (
	"fmt"
	_ "github.com/iamMarkchu/alpha/lib/orm"
	"github.com/iamMarkchu/alpha/router"
)

func main()  {
	if err := router.NewEngine().Run(); err != nil {
		fmt.Println("启动出现错误:", err.Error())
	}
}
