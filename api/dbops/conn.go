package dbops

import "database/sql"

var (
	dbConn *sql.DB		// 包内全局变量，用于长连接
	err error
)

func init() {
	// dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8"  此处严格区分大小写
	dbConn, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_video?charset=utf8")		// charset这些应作为配置项处理，这里简化，直接写在这里面
	if err != nil {
		panic(err.Error())
	}
}
