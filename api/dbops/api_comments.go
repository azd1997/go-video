package dbops

import (
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/utils"
)

func AddNewComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?,?,?,?) ")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {

	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
												INNER JOIN users ON comments.author_id = users.id
												WHERE comments.video_id = ? AND 
												comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
												ORDER BY comments.time DESC`)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	var id, authorName, content string
	var c *defs.Comment
	for rows.Next() {
		if err := rows.Scan(&id, &authorName, &content); err != nil {
			return res, nil
		}
		c = &defs.Comment{
			Id:      id,
			VideoId: vid,
			Author:  authorName,
			Content: content,
		}

		res = append(res, c)
	}

	return res, nil
}
