package dbops

import "testing"

var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("AddVideo: %v\n", err)
	}
	tempvid = video.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("GetVideo: %v\n", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempvid)
	if err != nil {
		t.Errorf("DeleteVideo: %v\n", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempvid)
	if err != nil || video != nil {
		t.Errorf("ReGetVideo: %v\n", err)
	}
}
