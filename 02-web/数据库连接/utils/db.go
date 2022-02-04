package utils

import (
	"database/sql"
	// 导入的这个mysql驱动不是给我们调用的, 而是给Open()函数调用的; sql.Open()建立数据库连接的时候, 需要使用mysql驱动, 所以我们必须先引入这个驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// 创建数据库连接池; 如果是本地的话, localhost:3306可以省略的
	Db, err = sql.Open("mysql", "root:123123@tcp(localhost:3306)/db_school")
	if err != nil {
		panic(err.Error())
	}
}
