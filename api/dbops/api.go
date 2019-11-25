package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	// 可以直接使用包内dbConn变量
}

func GetUserCredential(loginName string) (string, error) {
	// 可以直接使用包内dbConn变量
}