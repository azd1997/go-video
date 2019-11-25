package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	// 可以直接使用包内dbConn变量
	
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd)")		// prepare 预编译
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	// 可以直接使用包内dbConn变量

	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	
	stmDel.Exec(loginName, pwd)
	stmDel.Close()

	return nil
}