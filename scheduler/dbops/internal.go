package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// api -> videoid -> mysql
// dispatcher -> mysql -> videoid -> datachannel
// executor -> datachannel -> videoid -> delete videos


// 批量从数据库读 待删除 视频记录
func ReadVideoDeleteRecord(count int) ([]string, error) {

	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rc LIMIT ?")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("ReadVideoDeleteRecord: %v\n", err)
		return nil, err
	}

	var ids []string
	var id string
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func DelVideoDeleteRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rc WHERE video_id = ?")
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("DelVideoDeleteRecord: %v\n", err)
		return err
	}

	return nil
}