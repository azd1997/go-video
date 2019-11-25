package dbops

import (
	"database/sql"		// 通用接口

	_ "github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:123456!@#@tcp(localhost:3306)/go_video?charset=utf")		// charset这些应作为配置项处理，这里简化，直接写在这里面
	if err != nil {
		panic(err.Error())
	}

	return dbConn
}

func AddUserCredential(loginName string, pwd string) error {
	db := openConn()
}

func GetUserCredential(loginName string) (string, error) {
	db := openConn()
}