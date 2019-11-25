package dbops

import "testing"

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}



func testAddUser(t *testing.T) {
	err := AddUserCredential("eiger", "123456")
	if err != nil {
		t.Errorf("AddUser: %v\n", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("eiger")
	if pwd != "123456" || err != nil {
		t.Errorf("GetUser: %v\npwd = %s\n", err, pwd)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential("eiger", "123456")
	if err != nil {
		t.Errorf("DeleteUser: %v\n", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("eiger")
	if err != nil {
		t.Errorf("GetUser: %v\n", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}


