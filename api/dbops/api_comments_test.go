package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddComment", testAddNewComment)
	t.Run("ListComments", testListComments)
}

func testAddNewComment(t *testing.T) {
	vid, aid, content := "123456", 1, "I like this video"
	err := AddNewComment(vid, aid, content)
	if err != nil {
		t.Errorf("AddComment: %v\n", err)
	}
}

func testListComments(t *testing.T) {
	vid, from := "123456", 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("ListComments: %v\n", err)
	}

	for i, elem := range res {
		fmt.Printf("comment: %d, %v\n", i, elem)
	}
}

