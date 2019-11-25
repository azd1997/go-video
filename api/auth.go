package main

import (
	"github.com/azd1997/go-video/api/defs"
	"github.com/azd1997/go-video/api/session"
	"net/http"
)

// 自定义的HTTP头
var HEAD_FILE_SESSION = "X-Session-Id"
var HEAD_FILE_UNAME = "X-User-Name"

// 检查用户的session是否合法
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEAD_FILE_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {		// 表示session过期
		return false
	}

	r.Header.Add(HEAD_FILE_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEAD_FILE_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return false
	}

	return true
}
