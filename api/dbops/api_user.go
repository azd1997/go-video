package dbops

import (
	"database/sql"
	"github.com/azd1997/go-video/api/defs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	// 可以直接使用包内dbConn变量

	// stmt statement 是sql内的一个结构体，表示一个操作
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")		// prepare 预编译
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	// 可以直接使用包内dbConn变量

	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	defer stmtOut.Close()
	
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {		// 当查询到没有该行时一样会返回一个row去扫描给pwd，这样的情况下也会返回给err，
												// 但它其实是数据库没有该条记录我们应该额外在业务层面处理
		return "", err
	}
	
	return pwd, nil
}

func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	defer stmtDel.Close()
	
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	return nil
}

func GetUserInfo(loginName string) (*defs.UserInfo, error) {
	// 可以直接使用包内dbConn变量

	stmtOut, err := dbConn.Prepare("SELECT id FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer stmtOut.Close()

	var id int
	err = stmtOut.QueryRow(loginName).Scan(&id)
	if err != nil && err != sql.ErrNoRows {		// 当查询到没有该行时一样会返回一个row去扫描给pwd，这样的情况下也会返回给err，
		// 但它其实是数据库没有该条记录我们应该额外在业务层面处理
		return nil, err
	}

	return &defs.UserInfo{
		Id:       id,
		Username: loginName,
	}, nil
}