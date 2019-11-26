package dbops

import "log"

// 1. user -> api service -> delete video
// 2. api service -> scheduler -> write video delete record (写入数据库，说将要删哪个视频)
// 3. timer
// 4. timer -> runner -> read video delete record -> delete video from folder

func AddVideoDeleteRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rc (video_id) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeleteRecord: %v\n", err)
		return err
	}

	return nil
}
