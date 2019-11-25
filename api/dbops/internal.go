package dbops

import (
	"database/sql"
	"github.com/azd1997/go-video/api/defs"
	"log"
	"strconv"
	"sync"
)

// internal.go 存放session相关

func InsertSession(sid string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(sid, ttlstr, username)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	var ttl, username string
	err = stmtOut.QueryRow(sid).Scan(&ttl, username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var ttlint int64
	if ttlint, err = strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = ttlint
		ss.Username = username
	} else {
		return nil, err
	}

	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {

	m := &sync.Map{}

	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("RetrieveAllSessions: %s\n", err)
		return nil, err
	}
	defer stmtOut.Close()


	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("RetrieveAllSessions: %s\n", err)
		return nil, err
	}

	var id, ttlstr, username string
	var ttlint int64
	var ss *defs.SimpleSession
	for rows.Next() {
		if err = rows.Scan(&id, &ttlstr, &username); err != nil {
			log.Printf("RetrieveAllSessions: %v\n", err)
			break
		}

		if ttlint, err = strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss = &defs.SimpleSession{Username: username, TTL: ttlint}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl: %d\n", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("DeleteSession: %s\n", err)
		return err
	}

	if _, err = stmtDel.Exec(sid); err != nil {
		return err
	}

	return nil
}