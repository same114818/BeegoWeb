package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	// set default database
	_ = orm.RegisterDataBase("default", "mysql", "root:SJ114818@tcp(127.0.0.1:3306)/beego?charset=utf8", 30)

	_ = orm.RunSyncdb("default", false, false)
}
