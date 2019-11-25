package dbops

import "testing"

// api测试流程
// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)


func TestMain(m *testing.M) {
	// 1. 清除表数据
	clearTables()

	// 2. 运行测试流程
	m.Run()

	// 3. 清楚表数据
	clearTables()
}
