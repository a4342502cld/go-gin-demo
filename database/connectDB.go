package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	//_ "github.com/go-sql-driver/mysql"
)

func GetDataBase() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "a4342502"
	cfg.Net = "tcp"
	cfg.Addr = "mysql_5.7:3306"
	cfg.DBName = "gomysql"
	dsn := cfg.FormatDSN()

	log.Println(">>>> get database connection action start <<<<")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 返回数据库连接
	return db
}
