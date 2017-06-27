package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

//初始化数据库，包括设置数据库驱动，用户名，密码。返回一个数据库的连接

var session *sql.DB

func Connect() *sql.DB {
	return session
}

func Init() {
	mysqlUrl := beego.AppConfig.String("mysqlUrl")
	fmt.Println("mysql Init()", mysqlUrl)

	db, err := sql.Open("mysql", mysqlUrl)
	err1 := db.Ping()
	fmt.Println(err1)
	if err != nil {
		panic(err)
	} else {
		session = nil
		fmt.Println(err)
	}
	session = db
}
