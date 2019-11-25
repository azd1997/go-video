// session 会话，保存中间状态
// session与cookie的区别是session是服务端实现缓存的一种机制，而cookie是客户端的机制
// session需要有id，而session id会被缓存在cookie中

// 		客户端（人） -----------> api -----------> cache ------------> DB
//  			signin/register  		--write->      --write->
//		loggedin&return session id
//
//				session id  	-get session id->  -get session id->
//             logged in		<--session id--  	<--session id--

package session

import (
	"github.com/azd1997/go-video/api/dbops"
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/utils"
	"sync"
	"time"
)



var sessionMap *sync.Map	// 并发安全的map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	id, _ := utils.NewUUID()
	createTime := time.Now().UnixNano()/1000000		// 精确到 ms
	ttl := createTime + 30 * 60 * 1000				// 30 min 后失效

	ss := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, username)		// 插入到数据库中

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		createTime := nowInMilliSecond()
		if ss.(*defs.SimpleSession).TTL < createTime {		// 说明过期了
			deleteExpiredSession(sid)// 删除过期session
			return "", true
		}

		// 没过期把用户名返回
		return ss.(*defs.SimpleSession).Username, false
	} else {	// 已经过期且被删除，所以查不到
		return "", true
	}
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}


func nowInMilliSecond() int64 {
	return time.Now().UnixNano()/1000000
}