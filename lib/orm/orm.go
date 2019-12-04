package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engineGroup *xorm.EngineGroup
	err error
)

func init()  {
	masterDataSourceName := "root:root@tcp(127.0.0.1)/api_base"
	slave1DataSourceName := "root:root@tcp(127.0.0.1)/api_base"
	dataSourceNameSlice := []string{masterDataSourceName, slave1DataSourceName}
	if engineGroup, err = xorm.NewEngineGroup("mysql", dataSourceNameSlice); err != nil {
		fmt.Println("DB INIT出现错误", err)
	}
}

func NewEngine() *xorm.EngineGroup {
	return engineGroup
}
