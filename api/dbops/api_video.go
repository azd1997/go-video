package dbops

import (
	"database/sql"
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/utils"
	"time"
)

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")	// 格式的这个时间原点不能改

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
					(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCTime: ctime,
	}

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {

	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	var aid int
	var displayCTime, videoName string

	err = stmtOut.QueryRow(vid).Scan(&aid, &videoName, &displayCTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         videoName,
		DisplayCTime: displayCTime,
	}

	return res, nil
}

func DeleteVideo(vid string) error {

	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	return nil
}